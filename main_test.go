package main

import (
	"dataApi/Models"
	"database/sql"
	"fmt"
	"log"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// type repositoryDB struct {
// 	db *sql.DB
// }
var emp = &Models.Employee{
	Id:           uuid.New().String(),
	FirstName:    "Vinodh",
	MiddleName:   "K",
	LastName:     "Landa",
	Gender:       "Male",
	Salary:       555.55,
	DOB:          time.Now(),
	Email:        "vinod@gmail.com",
	Phone:        4634645,
	AddressLine1: "Lonsdale",
	AddressLine2: "street",
	State:        "vic",
	PostCode:     3000,
	TFN:          1354354,
	SuperBalance: 4645,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()

	fmt.Println("db status  is ", db.Stats().InUse)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestGetemployees(t *testing.T) {
	db, mock := NewMock()

	//sqlxdb := sql.NewDb(db, "sql mock")

	repo := &Models.Repository{db}

	defer repo.DB.Close()

	query := "SELECT (.+) FROM employee"

	rows := sqlmock.NewRows([]string{"Id", "FisrtName", "MiddleName",
		"LastName", "Gender", "Salary", "DOB", "Email", "Phone", "AddressLine1", "AddressLine2", "State", "PostCode", "TFN", "SuperBalance"}).
		AddRow("1", "Vinodh", "K", "Landa", "Male", 555.55, "1993-12-10", "vinod@gmail.com", "4634645", "Lonsdale", "street",
			"vic", "3000", 1354354, 4645)

	mock.ExpectQuery(query).WillReturnRows(rows)
	//	var c *gin.Context
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	gin.CreateTestContext(w)

	employees, err := repo.ExecuteGetAll(query)
	fmt.Println("in main test employees ", employees)
	assert.NotEmpty(t, employees)
	assert.Nil(t, err)

}

func TestCreateEmployeeRecordApi(t *testing.T) {

	db, mock := NewMock()
	repo := &Models.Repository{db}

	query := "INSERT INTO 'employee' (Id,FirstName,MiddleName,LastName,Gender,Salary,DOB,Email,Phone,AddressLine1,AddressLine2,State,PostCode,TFN,SuperBalance) VALUES (?, ?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	//querpart := "INSERT INTO 'employee'"
	//prep := mock.ExpectPrepare(query)
	// mock.ExpectExec().WillReturnResult()
	sql := regexp.QuoteMeta(query)
	//sqlmock.ExpectedExec()
	mock.MatchExpectationsInOrder(false)
	mock.ExpectExec(sql).WithArgs(true, 1).WillReturnResult(sqlmock.NewResult(0, 1))
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	gin.CreateTestContext(w)
	emp, err := repo.AddEmployee(Models.Employee{})
	assert.NotEmpty(t, emp)
	assert.NoError(t, err)

}

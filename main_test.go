package main

import (
	"dataApi/Models"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// type repository struct {
// 	db *sql.DB
// }

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()

	fmt.Println("db status  is ", db.Stats().InUse)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	return db, mock
}

func TestGetemployees(t *testing.T) {
	db, mock := NewMock()

	//sqlxdb := sql.NewDb(db, "sql mock")

	repo := &Models.Repository{db}

	defer func() {
		repo.Close()
	}()

	query := "SELECT * FROM employee"

	rows := sqlmock.NewRows([]string{"Id", "FisrtName", "MiddleName",
		"LastName", "Gender", "Salary", "DOB", "Email", "Phone", "AddressLine1", "AddressLine2", "State", "PostCode", "TFN", "SuperBalance"}).
		AddRow("1", "Vinodh", "K", "Landa", "Male", 555.55, "1993-12-10", "vinod@gmail.com", "4634645", "Lonsdale", "street",
			"vic", "3000", 1354354, 4645)

	mock.ExpectQuery(query).WillReturnRows(rows)

	// employees := Models.ExecuteQueryGetEmployees(query, c*&gin.Context)
	// fmt.Println(employees)
	// assert.NotEmpty(t, employees)
	// assert.NotNil(t, err)

}

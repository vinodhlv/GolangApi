package Models

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Repositoryinterface interface {
	Getemployees(c *gin.Context) []Employee
	GetemployeeById(c *gin.Context) []Employee
	CreateEmployeeRecordApi(c *gin.Context, emp *Employee) error
	DeleteEmployeeRecordApi(c *gin.Context)
	UpdateEmployeeRecordApi(c *gin.Context)
}

type Employee struct {
	Id           string    `json:"id"`
	FirstName    string    `json:"firstname"`
	MiddleName   string    `json:"middlename"`
	LastName     string    `json:"lastname"`
	Gender       string    `json:"gender"`
	Salary       float64   `json:"salary"`
	DOB          time.Time `json:"dob"`
	Email        string    `json:"email"`
	Phone        int       `json:"phone"`
	AddressLine1 string    `json:"address1"`
	AddressLine2 string    `json:"address2"`
	State        string    `json:"state"`
	PostCode     int       `json:"PostCode"`
	TFN          int       `json:"tfn"`
	SuperBalance float64   `json:"super"`
}

// type Repository struct {
// 	DB *sql.DB
// }

// Close attaches the provider and close the connection
// func (r *Repository) Close() {
// 	r.DB.Close()
// }

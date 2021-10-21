package Models

import (
	"time"
)

type Repositoryinterface interface {
	GetEmployees() ([]Employee, error)
	GetEmployeeById(id string) (Employee, error)
	AddEmployee(Employee) (Employee, error)
	DeleteEmployee(id string) error
	UpdateEmployee(Employee) error
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

package Models

import (
	"database/sql"
	"log"
)

type Repository struct {
	DB *sql.DB
}

// getemployees queries from the employee table .  (r *Repository)
func (r *Repository) GetEmployees() ([]Employee, error) {

	var employees []Employee
	query := "SELECT * FROM employee"
	rows, err := r.ExecuteGetAll(query)
	if err != nil {
		return employees, err
	}

	for rows.Next() {
		var emp Employee
		err = rows.Scan(&emp.Id, &emp.FirstName, &emp.MiddleName, &emp.LastName, &emp.Gender, &emp.Salary, &emp.DOB, &emp.Email, &emp.Phone, &emp.AddressLine1,
			&emp.AddressLine2, &emp.State, &emp.PostCode, &emp.TFN, &emp.SuperBalance)

		if err != nil {
			return employees, err
		}

		employees = append(employees, emp)
	}
	return employees, nil
}

// Loop through rows, returns the rows from the table.
func (r *Repository) ExecuteGetAll(query string) (*sql.Rows, error) {

	rows, err := r.DB.Query(query)
	if err != nil {
		return rows, err
	}

	return rows, err

}

// GetById retrive a record by id from the employee table .  (r *Repository)
func (r *Repository) GetEmployeeById(id string) (employee Employee, err error) {
	var emp Employee
	log.Println("In Model GetemployeeById Id value", id)

	query := "SELECT * FROM employee where id = ?"
	rows, err := r.ExecuteGetById(query, id)
	if err != nil {
		return emp, err
	}
	for rows.Next() {

		err = rows.Scan(&emp.Id, &emp.FirstName, &emp.MiddleName, &emp.LastName, &emp.Gender, &emp.Salary, &emp.DOB, &emp.Email, &emp.Phone, &emp.AddressLine1,
			&emp.AddressLine2, &emp.State, &emp.PostCode, &emp.TFN, &emp.SuperBalance)

	}
	return emp, err
}

// Loop through rows, returns the row of the id from the table.
func (r *Repository) ExecuteGetById(query string, id string) (*sql.Rows, error) {

	rows, err := r.DB.Query(query, id)
	log.Println("In Model ExecuteQueryemployeeById r value", r.DB.Stats())

	return rows, err

}

//Add row to the employee table
func (r *Repository) AddEmployee(emp Employee) (empl Employee, err error) {

	_, err = r.DB.Exec("INSERT INTO employee(Id,FirstName,MiddleName,LastName,Gender,Salary,DOB,Email,Phone,AddressLine1,AddressLine2,State,PostCode,TFN,SuperBalance) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		emp.Id, emp.FirstName, emp.MiddleName, emp.LastName, emp.Gender, emp.Salary, emp.DOB,
		emp.Email, emp.Phone, emp.AddressLine1, emp.AddressLine2, emp.State, emp.PostCode, emp.TFN,
		emp.SuperBalance)

	return emp, err

}

//Delete row from the employee table
func (r *Repository) DeleteEmployee(id string) error {

	_, err := r.DB.Exec("Delete from employee where id = ?", id)

	return err

}

//update the record from the employee table
func (r *Repository) UpdateEmployee(emp Employee) error {

	query := "Update employee set FirstName = ?,MiddleName = ?,LastName = ?,Gender = ?,Salary = ?,DOB = ?,Email = ?,Phone = ?,AddressLine1 = ?,AddressLine2 = ?,State = ?,PostCode = ?,TFN = ?,SuperBalance = ? where id = ?"
	_, err := r.DB.Exec(query, emp.FirstName, emp.MiddleName, emp.LastName, emp.Gender, emp.Salary, emp.DOB,
		emp.Email, emp.Phone, emp.AddressLine1, emp.AddressLine2, emp.State, emp.PostCode, emp.TFN,
		emp.SuperBalance, emp.Id)

	return err
}

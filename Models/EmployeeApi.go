package Models

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Repository struct {
	DB *sql.DB
}

// getemployees queries from the employee table .  (r *Repository)
func (r *Repository) Getemployees(c *gin.Context) []Employee {

	//var employees []Employee
	employees := make([]Employee, 0)
	query := "SELECT * FROM employee"
	rows := r.ExecuteQueryGetEmployees(query, c)

	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.Id, &emp.FirstName, &emp.MiddleName, &emp.LastName, &emp.Gender, &emp.Salary, &emp.DOB, &emp.Email, &emp.Phone, &emp.AddressLine1,
			&emp.AddressLine2, &emp.State, &emp.PostCode, &emp.TFN, &emp.SuperBalance)
		if c != nil {
			if err != nil {
				c.AbortWithStatus(http.StatusNotFound)
			}
			//c.IndentedJSON(http.StatusOK, emp) -- this call takes more bandwidth
			c.JSON(http.StatusOK, emp)
		}
		employees = append(employees, emp)
	}
	return employees
}

// Loop through rows, returns the rows from the table.
func (r *Repository) ExecuteQueryGetEmployees(query string, c *gin.Context) *sql.Rows {

	fmt.Println(query)
	rows, err := r.DB.Query(query)
	if err != nil {
		fmt.Println("Error after fetch", err)
		if c != nil {
			c.String(http.StatusNotFound, "Error while getting record ")
		}
	}

	return rows

}

// Loop through rows, returns the row of the id from the table.
func (r *Repository) ExecuteQueryGetEmployeeById(query string, c *gin.Context, id string) *sql.Rows {

	rows, err := r.DB.Query(query, id)
	log.Println("In Model ExecuteQueryemployeeById r value", r.DB.Stats())
	if err != nil {
		fmt.Println("Error after fetch", err)
		c.String(http.StatusNotFound, "Error while getting record ")
	}

	return rows

}

// GetemployeeById retrive a record by id from the employee table .  (r *Repository)
func (r *Repository) GetemployeeById(c *gin.Context) (employees []Employee) {

	id := c.Params.ByName("id")
	log.Println("In Model GetemployeeById Id value", id)
	query := "SELECT * FROM employee where id = ?"
	rows := r.ExecuteQueryGetEmployeeById(query, c, id)

	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.Id, &emp.FirstName, &emp.MiddleName, &emp.LastName, &emp.Gender, &emp.Salary, &emp.DOB, &emp.Email, &emp.Phone, &emp.AddressLine1,
			&emp.AddressLine2, &emp.State, &emp.PostCode, &emp.TFN, &emp.SuperBalance)
		if err != nil {
			c.String(http.StatusNotFound, "Error while getting record ")
		}
		//c.IndentedJSON(http.StatusOK, emp) -- this call takes more bandwidth
		c.JSON(http.StatusOK, emp)
	}
	return employees
}

func (r *Repository) CreateEmployeeRecordApi(c *gin.Context, emp *Employee) (err error) {

	c.BindJSON(&emp)
	_, err = r.DB.Exec("INSERT INTO employee(Id,FirstName,MiddleName,LastName,Gender,Salary,DOB,Email,Phone,AddressLine1,AddressLine2,State,PostCode,TFN,SuperBalance) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		emp.Id, emp.FirstName, emp.MiddleName, emp.LastName, emp.Gender, emp.Salary, emp.DOB,
		emp.Email, emp.Phone, emp.AddressLine1, emp.AddressLine2, emp.State, emp.PostCode, emp.TFN,
		emp.SuperBalance)

	if err != nil {
		c.String(http.StatusInternalServerError, "Error while inserting record ")
	}
	c.JSON(http.StatusOK, emp)
	return err

}

func (r *Repository) DeleteEmployeeRecordApi(c *gin.Context) {

	fmt.Println("In the Delete Api")
	var emp Employee
	c.BindJSON(&emp)
	id := c.Params.ByName("id")
	fmt.Println("The id value is ", id)
	rows, err := r.DB.Exec("Delete from employee where id = ?", id)
	fmt.Println("The rows and err value are", rows, err)

	if err != nil {
		c.String(http.StatusInternalServerError, "Error while deleting record ")
	} else {
		c.JSON(http.StatusOK, gin.H{"id #" + id: "deleted"})
		rows.RowsAffected()
	}

}

func (r *Repository) UpdateEmployeeRecordApi(c *gin.Context) {

	fmt.Println("In the Update Api")
	var emp Employee
	c.BindJSON(&emp)
	id := c.Params.ByName("id")
	fmt.Println("The id value is ", id)
	query := "Update employee set FirstName = ?,MiddleName = ?,LastName = ?,Gender = ?,Salary = ?,DOB = ?,Email = ?,Phone = ?,AddressLine1 = ?,AddressLine2 = ?,State = ?,PostCode = ?,TFN = ?,SuperBalance = ? where id = ?"
	rows, err := r.DB.Exec(query, emp.FirstName, emp.MiddleName, emp.LastName, emp.Gender, emp.Salary, emp.DOB,
		emp.Email, emp.Phone, emp.AddressLine1, emp.AddressLine2, emp.State, emp.PostCode, emp.TFN,
		emp.SuperBalance, id)
	fmt.Println("The rows and err value are", rows, err)

	if err != nil {
		c.String(http.StatusInternalServerError, "Error while updating record ")
	} else {
		//count,err:=rows.RowsAffected()
		c.JSON(http.StatusOK, gin.H{"id #" + id: "updated"})
	}
}

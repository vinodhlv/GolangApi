package Models

import (
	"dataApi/Config"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getemployees queries from the employee table .  (r *Repository)
func Getemployees(c *gin.Context) []Employee {

	var employees []Employee

	query := "SELECT * FROM employee"
	rows := ExecuteQueryGetEmployees(query, c)

	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.Id, &emp.FirstName, &emp.MiddleName, &emp.LastName, &emp.Gender, &emp.Salary, &emp.DOB, &emp.Email, &emp.Phone, &emp.AddressLine1,
			&emp.AddressLine2, &emp.State, &emp.PostCode, &emp.TFN, &emp.SuperBalance)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		}
		//c.IndentedJSON(http.StatusOK, emp) -- this call takes more bandwidth
		c.JSON(http.StatusOK, emp)
	}
	return employees
}

// Loop through rows, returns the rows from the table.
func ExecuteQueryGetEmployees(query string, c *gin.Context) *sql.Rows {

	rows, err := Config.DB.Query(query)

	if err != nil {
		fmt.Println("Error after fetch", err)
		c.AbortWithStatus(http.StatusNotFound)

	}

	return rows

}

// Loop through rows, returns the row of the id from the table.
func ExecuteQueryGetEmployeeById(query string, c *gin.Context, id string) *sql.Rows {

	rows, err := Config.DB.Query(query, id)

	if err != nil {
		fmt.Println("Error after fetch", err)
		c.AbortWithStatus(http.StatusNotFound)

	}

	return rows

}

// GetemployeeById retrive a record by id from the employee table .  (r *Repository)
func GetemployeeById(c *gin.Context) []Employee {

	var employees []Employee
	id := c.Params.ByName("id")
	query := "SELECT * FROM employee where id = ?"
	rows := ExecuteQueryGetEmployeeById(query, c, id)

	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.Id, &emp.FirstName, &emp.MiddleName, &emp.LastName, &emp.Gender, &emp.Salary, &emp.DOB, &emp.Email, &emp.Phone, &emp.AddressLine1,
			&emp.AddressLine2, &emp.State, &emp.PostCode, &emp.TFN, &emp.SuperBalance)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		}
		//c.IndentedJSON(http.StatusOK, emp) -- this call takes more bandwidth
		c.JSON(http.StatusOK, emp)
	}
	return employees
}

func CreateEmployeeRecordApi(c *gin.Context) {

	var emp Employee
	c.BindJSON(&emp)
	_, err := Config.DB.Exec("INSERT INTO employee(Id,FirstName,MiddleName,LastName,Gender,Salary,DOB,Email,Phone,AddressLine1,AddressLine2,State,PostCode,TFN,SuperBalance) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		emp.Id, emp.FirstName, emp.MiddleName, emp.LastName, emp.Gender, emp.Salary, emp.DOB,
		emp.Email, emp.Phone, emp.AddressLine1, emp.AddressLine2, emp.State, emp.PostCode, emp.TFN,
		emp.SuperBalance)

	if err != nil {
		c.String(http.StatusInternalServerError, "Error while inserting record ")
	}
	c.JSON(http.StatusOK, emp)

}

func DeleteEmployeeRecordApi(c *gin.Context) {

	fmt.Println("In the Delete Api")
	var emp Employee
	c.BindJSON(&emp)
	id := c.Params.ByName("id")
	fmt.Println("The id value is ", id)
	rows, err := Config.DB.Exec("Delete from employee where id = ?", id)
	fmt.Println("The rows and err value are", rows, err)

	if err != nil {
		c.String(http.StatusInternalServerError, "Error while deleting record ")
	} else {
		c.JSON(http.StatusOK, gin.H{"id #" + id: "deleted"})
		rows.RowsAffected()
	}

}

func UpdateEmployeeRecordApi(c *gin.Context) {

	fmt.Println("In the Update Api")
	var emp Employee
	c.BindJSON(&emp)
	id := c.Params.ByName("id")
	fmt.Println("The id value is ", id)
	query := "Update employee set FirstName = ?,MiddleName = ?,LastName = ?,Gender = ?,Salary = ?,DOB = ?,Email = ?,Phone = ?,AddressLine1 = ?,AddressLine2 = ?,State = ?,PostCode = ?,TFN = ?,SuperBalance = ? where id = ?"
	rows, err := Config.DB.Exec(query, emp.FirstName, emp.MiddleName, emp.LastName, emp.Gender, emp.Salary, emp.DOB,
		emp.Email, emp.Phone, emp.AddressLine1, emp.AddressLine2, emp.State, emp.PostCode, emp.TFN,
		emp.SuperBalance, id)
	fmt.Println("The rows and err value are", rows, err)

	if err != nil {
		c.String(http.StatusInternalServerError, "Error while updating record ")
	} else {
		c.JSON(200, gin.H{"id #" + id: "updated"})
		rows.RowsAffected()
	}
}

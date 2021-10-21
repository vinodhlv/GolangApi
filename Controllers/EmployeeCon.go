package Controllers

import (
	"dataApi/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context, repo Models.Repositoryinterface) {

	employees, err := repo.GetEmployees()
	if err != nil {
		c.JSON(http.StatusNotFound, "Error while getting records ")
	}
	c.JSON(http.StatusOK, employees)
}

func GetById(c *gin.Context, repo Models.Repositoryinterface) {
	id := c.Params.ByName("id")
	employee, err := repo.GetEmployeeById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "Error while getting record ")
		//c.IndentedJSON(http.StatusOK, emp) -- this call takes more bandwidth
	}

	c.JSON(http.StatusOK, employee)

}

func Add(c *gin.Context, repo Models.Repositoryinterface) {
	var employee Models.Employee
	c.BindJSON(&employee)
	emp, err := repo.AddEmployee(employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error while inserting record ")
	}

	c.JSON(http.StatusCreated, emp)
}

func Delete(c *gin.Context, repo Models.Repositoryinterface) {
	id := c.Params.ByName("id")
	var emp Models.Employee
	c.BindJSON(&emp)
	err := repo.DeleteEmployee(id)

	if err != nil {
		c.String(http.StatusInternalServerError, "Error while deleting record ")
	} else {
		c.JSON(http.StatusOK, gin.H{"id #" + id: "deleted"})
	}
}

func Save(c *gin.Context, repo Models.Repositoryinterface) {

	id := c.Params.ByName("id")
	var emp Models.Employee
	c.BindJSON(&emp)
	emp.Id = id
	err := repo.UpdateEmployee(emp)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error while updating record ")
	} else {
		c.JSON(http.StatusOK, gin.H{"id #" + id: "updated"})
	}
}

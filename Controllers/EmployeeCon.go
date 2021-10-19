package Controllers

import (
	"dataApi/Models"

	"github.com/gin-gonic/gin"
)

func GetEmployeesNavigate(c *gin.Context) {

	// var DB *sqlx.DB

	// repo := &Models.Repository{DB}

	// repo.Getemployees()

	Models.Getemployees(c)

}

func GetEmployeeByIdNavigate(c *gin.Context) {

	Models.GetemployeeById(c)

}

func CreateEmployeeRecord(c *gin.Context) {

	Models.CreateEmployeeRecordApi(c)
}

func DeleteEmployeeRecord(c *gin.Context) {

	Models.DeleteEmployeeRecordApi(c)
}

func UpdateEmployeeRecord(c *gin.Context) {

	Models.UpdateEmployeeRecordApi(c)
}

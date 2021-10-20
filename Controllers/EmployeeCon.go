package Controllers

import (
	"dataApi/Config"
	"dataApi/Models"

	"github.com/gin-gonic/gin"
)

var repoInterf Models.Repositoryinterface
var dbcon Models.Repository

func GetEmployeesNavigate(c *gin.Context) {

	dbcon.DB = Config.GetDB()

	repoInterf = &dbcon

	repoInterf.Getemployees(c)

}

func GetEmployeeByIdNavigate(c *gin.Context) {

	dbcon.DB = Config.GetDB()

	repoInterf = &dbcon
	repoInterf.GetemployeeById(c)

	// DB := Config.GetDB()
	// repo := &Models.Repository{DB}
	// repo.GetemployeeById(c)
}

func CreateEmployeeRecord(c *gin.Context) {
	dbcon.DB = Config.GetDB()

	repoInterf = &dbcon
	var emp *Models.Employee

	repoInterf.CreateEmployeeRecordApi(c, emp)
}

func DeleteEmployeeRecord(c *gin.Context) {
	dbcon.DB = Config.GetDB()

	repoInterf = &dbcon

	repoInterf.DeleteEmployeeRecordApi(c)
}

func UpdateEmployeeRecord(c *gin.Context) {
	dbcon.DB = Config.GetDB()

	repoInterf = &dbcon

	repoInterf.UpdateEmployeeRecordApi(c)
}

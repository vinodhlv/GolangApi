package Controllers

import (
	"dataApi/Config"
	"dataApi/Models"
	"database/sql"

	"github.com/gin-gonic/gin"
)

var repoInterf Models.Repositoryinterface
var dbcon Models.Repository
var DB *sql.DB

func getRepo() *Models.Repository {
	dbcon.DB = Config.GetDB()
	return &dbcon
}

func GetEmployeesNavigate(c *gin.Context) {

	//dbcon.DB = Config.GetDB()

	repoInterf = getRepo()

	repoInterf.Getemployees(c)

}

func GetEmployeeByIdNavigate(c *gin.Context) {

	dbcon.DB = Config.GetDB()

	repoInterf = getRepo()

	repoInterf.GetemployeeById(c)

	// DB := Config.GetDB()
	// repo := &Models.Repository{DB}
	// repo.GetemployeeById(c)
}

func CreateEmployeeRecord(c *gin.Context) {

	repoInterf = getRepo()

	var emp *Models.Employee

	repoInterf.CreateEmployeeRecordApi(c, emp)
}

func DeleteEmployeeRecord(c *gin.Context) {

	repoInterf = getRepo()

	repoInterf.DeleteEmployeeRecordApi(c)
}

func UpdateEmployeeRecord(c *gin.Context) {

	repoInterf = getRepo()

	repoInterf.UpdateEmployeeRecordApi(c)
}

package Routes

import (
	"dataApi/Controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	route := gin.Default()

	group1 := route.Group("/employeeApi")
	{
		group1.GET("Getemployees", Controllers.GetEmployeesNavigate)
		group1.POST("Insertemployee", Controllers.CreateEmployeeRecord)
		group1.DELETE("Deleteemployee/:id", Controllers.DeleteEmployeeRecord)
		group1.PUT("Updateemployee/:id", Controllers.UpdateEmployeeRecord)
		group1.GET("Getemployee/:id", Controllers.GetEmployeeByIdNavigate)
	}

	return route
}

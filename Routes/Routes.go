package Routes

import (
	"dataApi/Config"
	"dataApi/Controllers"
	"dataApi/Models"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	var repoInterf Models.Repositoryinterface
	var dbcon Models.Repository

	DB := Config.GetDB()

	dbcon = Models.Repository{DB}
	repoInterf = &dbcon

	route := gin.Default()
	route.GET("/employee/:id", func(c *gin.Context) { Controllers.GetById(c, repoInterf) })
	route.GET("/employees", func(c *gin.Context) { Controllers.GetAll(c, repoInterf) })
	route.POST("/employee", func(c *gin.Context) { Controllers.Add(c, repoInterf) })
	route.DELETE("/employee/:id", func(c *gin.Context) { Controllers.Delete(c, repoInterf) })
	route.PUT("/employee/:id", func(c *gin.Context) { Controllers.Save(c, repoInterf) })

	return route
}

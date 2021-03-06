package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		debit := main.Group("debit")
		{
			debit.GET("/", controllers.ShowAllDebits)
			debit.GET("/:id", controllers.ShowDebit)
			debit.POST("/", controllers.CreateBook)
		}
	}
	return router
}

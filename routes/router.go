package routes

import (
	"todo-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/todos", controllers.CreateTodoHandler)
	router.GET("/todos", controllers.GetTodosHandler)
	router.PUT("/todos/:id", controllers.UpdateTodoHandler)
	router.DELETE("/todos/:id", controllers.DeleteTodoHandler)

	return router
}

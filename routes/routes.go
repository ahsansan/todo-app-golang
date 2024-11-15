package routes

import (
	"todo-app/controllers"
	"todo-app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		todo := api.Group("/todos").Use(middlewares.AuthMiddleware())
		{
			todo.POST("/", controllers.CreateTodo)
			todo.GET("/", controllers.GetTodos)
		}
	}
}
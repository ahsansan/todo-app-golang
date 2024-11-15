package main

import (
	"todo-app/config"
	"todo-app/models"
	"todo-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect to database
	config.ConnectDatabase()

	// Auto migrate models
	config.DB.AutoMigrate(&models.User{}, &models.Todo{})

	// Setup routes
	routes.SetupRoutes(router)

	// Run server
	router.Run(":8080")
}
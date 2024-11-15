package controllers

import (
	"net/http"
	"todo-app/config"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

// Create Todo
func CreateTodo(c *gin.Context) {
	// Ambil user ID dari context
	userIdInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Konversi user ID ke uint
	userId, ok := userIdInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	var todo models.Todo

	// Validasi payload
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set CreatedBy dari user ID
	todo.CreatedBy = userId
	todo.IsCompleted = false

	// Simpan todo ke database
	if err := config.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// Get User's Todos
func GetTodos(c *gin.Context) {
	userId, _ := c.Get("user_id")
	var todos []models.Todo

	config.DB.Where("created_by = ?", userId).Find(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

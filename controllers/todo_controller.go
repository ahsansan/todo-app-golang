package controllers

import (
	"net/http"
	"todo-app/config"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

// Create Todo
func CreateTodo(c *gin.Context) {
	userId, _ := c.Get("user_id")

	var todo models.Todo

	// Validasi payload
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set CreatedBy dari user ID
	todo.CreatedBy = userId.(uint)
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

// Complete Todo
func CompleteTodo(c *gin.Context) {
	userId, _ := c.Get("user_id")
	todoID := c.Param("id") // Ambil ID Todo dari parameter URL

	// Struktur untuk input request body
	var input struct {
		IsCompleted bool `json:"is_completed"`
	}

	// Validasi input dari request body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cari todo berdasarkan ID dan milik user yang sedang login
	var todo models.Todo
	if err := config.DB.Where("id = ? AND created_by = ?", todoID, userId).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found or not owned by you"})
		return
	}

	// Perbarui status is_completed
	todo.IsCompleted = input.IsCompleted
	if err := config.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	// Berikan respon sukses
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// Delete Todo
func DeleteTodo(c *gin.Context) {
	userId, _ := c.Get("user_id")
	todoID := c.Param("id") // Ambil ID Todo dari parameter URL

	// Cari todo berdasarkan ID dan milik user yang sedang login
	var todo models.Todo
	if err := config.DB.Where("id = ? AND created_by = ?", todoID, userId).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found or not owned by you"})
		return
	}

	// Hapus todo dari database
	if err := config.DB.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	// Berikan respon sukses
	c.JSON(http.StatusOK, gin.H{"message": "Todo successfully deleted"})
}
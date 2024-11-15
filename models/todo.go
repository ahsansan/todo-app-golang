package models

import "time"

type Todo struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `json:"title"`
	Content string    `json:"content"`
	IsCompleted bool      `json:"is_completed"`
	CreatedBy   uint      `json:"created_by"`
	// User        User      `gorm:"foreignKey:CreatedBy"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
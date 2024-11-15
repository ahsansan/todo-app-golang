package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Email string
	FullName string
	Password string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
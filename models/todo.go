package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title     string
	Content   string
	IsCompleted bool
	CreatedBy uint
}
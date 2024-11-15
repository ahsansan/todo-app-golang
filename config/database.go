package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// dsn := "root:password@tcp(127.0.0.1:3306)/todo_app_go?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root@tcp(127.0.0.1:3306)/todo_app_go?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = database
}
package database

import (
	"go-graphql/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("❌ Failed to connect database")
	}
	println("✅ Connected to database")
	db.AutoMigrate(&models.Book{})
	DB = db
}

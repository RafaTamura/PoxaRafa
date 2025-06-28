package database

import (
	"log"
	"poxa_rafa/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("poxarafa.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error during connect:", err)
	}
	DB.AutoMigrate(&models.Post{})
	log.Println("Database connected successfully")
}

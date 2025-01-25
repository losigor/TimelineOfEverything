package storage

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("DB could not load .env file", err)
	}
	dsn := os.Getenv("DSN")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not be open DB", err)
	}

	DB.AutoMigrate(&User{})
}

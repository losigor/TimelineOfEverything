package storage

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Event struct {
	ID   int    `json: "id"`
	Name string `json: "name"`
	//TimeDate time.Time `json: "timedate` // проверить Time в SQL
	TimeDate string `json: "timedate"`
	Text     string `json: "text"`
}

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("DB could not load .env file")
	}
	dsn := os.Getenv("DSN")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not open DB: ", err)
	}

	DB.AutoMigrate(&Event{})
}

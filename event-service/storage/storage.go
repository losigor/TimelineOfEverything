package storage

import (
	"log"

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

// подбить к конфиг-файлу или .енв
var dsn = "host=localhost user=postgres password=timeline dbname=postgres port=5432 sslmode=disable"
var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not open DB: ", err)
	}

	DB.AutoMigrate(&Event{})
}

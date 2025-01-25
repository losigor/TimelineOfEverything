package main

import (
	"os"
	"user-service/internal/handlers"
	"user-service/storage"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal("Could not load .env file")
	}

	storage.InitDB()

	e.POST("/register", handlers.RegisterUser)

	appAdress := os.Getenv("APP_ADRESS")
	e.Logger.Fatal(e.Start(appAdress))
}

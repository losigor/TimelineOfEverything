package main

import (
	"event-service/internal/handlers"
	"event-service/storage"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	if err := godotenv.Load(); err != nil {
		e.Logger.Fatal("Error loading .env file")
	}

	storage.InitDB()

	e.POST("/events", handlers.PostEvent)
	e.GET("/events/:id", handlers.GetEventByID)
	e.GET("/events", handlers.GetEvents)
	e.PATCH("/events/:id", handlers.PatchEvent)
	e.DELETE("/events/:id", handlers.DeleteEvent)

	appAdress := os.Getenv("APP_ADRESS")
	e.Logger.Fatal(e.Start(appAdress))

}

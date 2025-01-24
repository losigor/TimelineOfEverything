package main

import (
	"Timeline/event-service/internal/handlers"
	"Timeline/event-service/storage"

	"github.com/labstack/echo/v4"
)

func main() {

	storage.InitDB()

	e := echo.New()
	e.POST("/events", handlers.PostEvent)
	e.GET("/events/:id", handlers.GetEventByID)
	e.GET("/events", handlers.GetEvents)
	e.PATCH("/events/:id", handlers.PatchEvent)
	e.DELETE("/events/:id", handlers.DeleteEvent)

	e.Start(":8083")

}

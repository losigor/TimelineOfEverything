package main

import (
	"Timeline/event-service/internal/handlers"
	"Timeline/event-service/storage"

	"github.com/labstack/echo/v4"
)

func main() {

	storage.InitDB()

	e := echo.New()
	e.POST("/events", handlers.PostHandler)
	e.GET("/events/:id", handlers.GetHandlerByID)
	e.GET("/events", handlers.GetHandler)
	e.PATCH("/events/:id", handlers.PatchHandler)
	e.DELETE("/events/:id", handlers.DeleteHandler)

	e.Start(":8083")

}

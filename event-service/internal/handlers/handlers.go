package handlers

import (
	. "event-service/storage"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  string `json: "status"`
	Message string `json: "message"`
}

func PostEvent(c echo.Context) error {
	var event Event

	if err := c.Bind(&event); err != nil {
		return BadRequest(c, "Could not bind event")
	}

	if err := DB.Where("text = ?", event.Name).First(&Event{}); err == nil {
		return BadRequest(c, "Event with this name is already exists")
	}

	if err := DB.Create(&event).Error; err != nil {
		return BadRequest(c, "Could not add to DB")
	}

	return Success(c, "Event was added")
}

func GetEventByID(c echo.Context) error {

	var event Event
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)

	if err != nil {
		return BadRequest(c, "Bad ID. You should use only numbers")
	}

	if err = DB.Model(&Event{}).Where("id = ?", id).Find(&event).Error; err != nil || event.ID == 0 {
		return BadRequest(c, "Could not find event with this ID")
	}

	return c.JSON(http.StatusOK, &event)
}

func GetEvents(c echo.Context) error {

	var events []Event

	if err := DB.Limit(10).Find(&events).Error; err != nil {
		return BadRequest(c, "Could not find event with this ID")
	}

	return c.JSON(http.StatusOK, &events)
}

func PatchEvent(c echo.Context) error {

	ParamID := c.Param("id")
	id, err := strconv.Atoi(ParamID)
	if err != nil {
		return BadRequest(c, "Bad ID. You should use onle numbers")
	}

	fmt.Println(id)

	var updatedEvent Event
	if err = c.Bind(&updatedEvent); err != nil {
		return BadRequest(c, "Invalid input(Bind error)")
	}

	err = DB.Model(&Event{}).Where("id = ?", id).Updates(Event{
		Name: updatedEvent.Name, TimeDate: updatedEvent.TimeDate, Text: updatedEvent.Text}).Error

	if err != nil {
		return BadRequest(c, "Could not update the event(DB update error)")
	}

	return Success(c, "Event was successfully updated")
}

func DeleteEvent(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return BadRequest(c, "Bad ID. You should use only numbers")
	}

	if err = DB.Delete(&Event{}, id).Error; err != nil {
		return BadRequest(c, "Could not delete the event(DB delete error)")
	}

	return Success(c, "Event was successfully deleted")
}

func BadRequest(c echo.Context, s string) error {
	return c.JSON(http.StatusBadRequest, Response{
		Status:  "Error",
		Message: s,
	})
}

func Success(c echo.Context, s string) error {
	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: s,
	})
}

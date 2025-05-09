package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateTask(c echo.Context) error {
	var body RequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid data"})
	}

	task := Task{
		ID:          uuid.New(),
		Title:       body.Task,
		IsCompleted: false,
	}

	tasks = append(tasks, task)
	return c.JSON(http.StatusCreated, task)
}

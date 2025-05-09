package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RequestBody struct {
	Task string `json:"task"`
}

func UpdateTitleTaskById(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "invalid_id",
			"message": fmt.Sprintf("Неверный формат id: %v", err),
		})
	}

	var body RequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "Invalid request data",
			"message": "Проверьте введенные данные",
		})
	}

	for index, task := range tasks {
		if task.ID == id {
			tasks[index].Title = body.Task
			return c.JSON(http.StatusOK, tasks[index])
		}
	}

	return c.JSON(http.StatusNotFound, echo.Map{
		"error":   "not_found",
		"message": fmt.Sprintf("Task с id %s не найден", id),
	})
}

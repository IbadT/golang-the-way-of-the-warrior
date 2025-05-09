package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func DeleteTaskById(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "invalid_id",
			"message": fmt.Sprintf("Неверный формат id: %v", err),
		})
	}

	for index, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:index], tasks[index+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.JSON(http.StatusNotFound, echo.Map{
		"error":   "not_found",
		"message": fmt.Sprintf("Task с id %s не найден", id),
	})
}

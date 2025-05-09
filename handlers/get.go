package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Task struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	IsCompleted bool      `json:"is_completed"`
}

var tasks = []Task{}

func GetTasks(c echo.Context) error {
	isCompletedStr := c.QueryParam("is_completed")

	if isCompletedStr != "" {
		isCompleted, err := strconv.ParseBool(isCompletedStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error":   "invalid_boolean",
				"message": fmt.Sprintf("Неверное значение параметра is_completed: %s", isCompletedStr),
			})
		}

		filteredTasks := make([]Task, 0, len(tasks))
		for _, task := range tasks {
			if task.IsCompleted == isCompleted {
				filteredTasks = append(filteredTasks, task)
			}
		}

		return c.JSON(http.StatusOK, filteredTasks)
	}

	return c.JSON(http.StatusOK, tasks)

}

func GetTaskById(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request id")
	}

	for _, task := range tasks {
		if task.ID == id {
			return c.JSON(http.StatusOK, task)
		}
	}

	return c.JSON(http.StatusNotFound, echo.Map{
		"error":   "not_found",
		"message": fmt.Sprintf("Task с id %s не найден", id),
	})
}

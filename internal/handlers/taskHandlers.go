package handlers

import (
	"net/http"

	taskservice "github.com/IbadT/golang-the-way-of-the-warrior.git/internal/taskService"
	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service taskservice.TaskService
}

func NewTaskHandler(s taskservice.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

// GET
func (h *TaskHandler) GetTasks(c echo.Context) error {
	queryParams := c.QueryParams()

	if len(queryParams) > 0 && queryParams.Get("is_done") == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "Invalid query param",
			"message": "Нужно передать только параметр is_done",
		})
	}

	isDoneQuery := c.QueryParam("is_done")
	tasks, err := h.service.GetTasks(isDoneQuery)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Could not get Tasks",
		})
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTaskById(c echo.Context) error {
	idStr := c.Param("id")

	task, err := h.service.GetTaskById(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Could not find task",
		})
	}
	return c.JSON(http.StatusOK, task)
}

// POST
func (h *TaskHandler) CreateTask(c echo.Context) error {
	var body taskservice.RequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid data",
		})
	}

	task, err := h.service.CreateTask(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Could not create task",
		})
	}

	return c.JSON(http.StatusCreated, task)
}

// PATCH
func (h *TaskHandler) UpdateTaskCompletedById(c echo.Context) error {
	idStr := c.Param("id")

	task, err := h.service.UpdateTaskCompletedById(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Could not update task",
		})
	}

	return c.JSON(http.StatusOK, task)
}

// PUT
func (h *TaskHandler) UpdateTitleTaskById(c echo.Context) error {
	idStr := c.Param("id")

	var body taskservice.RequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "Invalid request data",
			"message": "Проверьте введенные данные",
		})
	}

	task, err := h.service.UpdateTitleTaskById(idStr, body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Could not update task title",
		})
	}

	return c.JSON(http.StatusOK, task)
}

// DELETE
func (h *TaskHandler) DeleteTaskById(c echo.Context) error {
	idStr := c.Param("id")

	if err := h.service.DeleteTaskById(idStr); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Could not delete task",
		})
	}

	return c.NoContent(http.StatusNoContent)
}

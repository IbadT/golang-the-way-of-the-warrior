package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// 2. Добавить глобальную переменную `task`
// 3. Добавить `POST` handler, который будет принимать json с полем `task` и записывать его содержимое в нашу переменную.
// 4. Обновить `GET` handler, чтобы он возвращал “hello, `task` ”

type Task struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	IsCompleted bool      `json:"is_completed"`
}

var tasks = []Task{}

func main() {
	e := echo.New()

	// xh http://localhost:8080/tasks
	e.GET("/tasks", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello, 'task'")
	})

	// xh POST http://localhost:8080/tasks title="Новая задача"
	e.POST("/tasks", func(c echo.Context) error {
		var task Task
		if err := c.Bind(&task); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid data"})
		}
		if task.ID == uuid.Nil {
			task.ID = uuid.New()
		}
		tasks = append(tasks, Task{
			ID:          task.ID,
			Title:       task.Title,
			IsCompleted: false,
		})
		return c.JSON(http.StatusCreated, task)
	})

	e.Start(":8080")
}

package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// 2. Добавить глобальную переменную `task`
// 3. Добавить `POST` handler, который будет принимать json с полем `task` и записывать его содержимое в нашу переменную.
// 4. Обновить `GET` handler, чтобы он возвращал “hello, `task` ”

type RequestBody struct {
	Task string `json:"task"`
}

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

	// xh POST http://localhost:8080/tasks task="Новая задача"
	e.POST("/tasks", func(c echo.Context) error {
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
	})

	e.Use(middleware.Logger())
	e.Start(":8080")
}

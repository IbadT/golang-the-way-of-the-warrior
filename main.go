package main

import (
	"github.com/IbadT/golang-the-way-of-the-warrior.git/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// 1. Реализовать `PATCH` ручку, которая будет обновлять `task` *(Саму задачу или ее статус)* по `ID`
// 2. Реализовать `DELETE` ручку, которая будет удалять `task` по `ID`

func main() {
	e := echo.New()

	// xh http://localhost:8080/tasks
	e.GET("/tasks", handlers.GetTasks)
	// xh http://localhost:8080/tasks/{id}
	e.GET("/tasks/:id", handlers.GetTaskById)

	// xh POST http://localhost:8080/tasks task="Новая задача"
	e.POST("/tasks", handlers.CreateTask)

	// xh patch http://localhost:8080/tasks/{id}
	e.PATCH("/tasks/:id", handlers.UpdateTaskCompletedById)
	// xh put http://localhost:8080/tasks/{id} task="Обновленная задача"
	e.PUT("/tasks/:id", handlers.UpdateTitleTaskById)

	// xh delete http://localhost:8080/tasks/{id}
	e.DELETE("/tasks/:id", handlers.DeleteTaskById)

	e.Use(middleware.Logger())
	e.Start(":8080")
}

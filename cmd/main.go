package main

import (
	"log"

	"github.com/IbadT/golang-the-way-of-the-warrior.git/internal/db"
	"github.com/IbadT/golang-the-way-of-the-warrior.git/internal/handlers"
	taskservice "github.com/IbadT/golang-the-way-of-the-warrior.git/internal/taskService"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	taskRepo := taskservice.NewTaskRepository(database)
	taskService := taskservice.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)

	e := echo.New()

	e.Use(middleware.Logger())

	// xh http://localhost:8080/tasks
	// xh http://localhost:8080/tasks?is_done=true
	e.GET("/tasks", taskHandlers.GetTasks)
	// xh http://localhost:8080/tasks/{id}
	e.GET("/tasks/:id", taskHandlers.GetTaskById)

	// xh POST http://localhost:8080/tasks task="Новая задача"
	e.POST("/tasks", taskHandlers.CreateTask)

	// xh patch http://localhost:8080/tasks/{id}
	e.PATCH("/tasks/:id", taskHandlers.UpdateTaskCompletedById)
	// xh put http://localhost:8080/tasks/{id} task="Обновленная задача"
	e.PUT("/tasks/:id", taskHandlers.UpdateTitleTaskById)

	// xh delete http://localhost:8080/tasks/{id}
	e.DELETE("/tasks/:id", taskHandlers.DeleteTaskById)

	e.Start(":8080")
}

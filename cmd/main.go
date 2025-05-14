package main

import (
	"log"

	"github.com/IbadT/golang-the-way-of-the-warrior.git/internal/db"
	"github.com/IbadT/golang-the-way-of-the-warrior.git/internal/handlers"
	taskservice "github.com/IbadT/golang-the-way-of-the-warrior.git/internal/taskService"
	"github.com/IbadT/golang-the-way-of-the-warrior.git/internal/web/tasks"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	repo := taskservice.NewTaskRepository(database)
	service := taskservice.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	e.Start(":8080")
}

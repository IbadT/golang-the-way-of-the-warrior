package main

import (
	"log"

	"github.com/IbadT/golang-the-way-of-the-warrior.git/internal/db"
	"github.com/IbadT/golang-the-way-of-the-warrior.git/internal/handlers"
	taskservice "github.com/IbadT/golang-the-way-of-the-warrior.git/internal/taskService"
	userservice "github.com/IbadT/golang-the-way-of-the-warrior.git/internal/userService"
	"github.com/IbadT/golang-the-way-of-the-warrior.git/internal/web/tasks"
	"github.com/IbadT/golang-the-way-of-the-warrior.git/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Println("Ошибка подключения к базе данных:", err)
		log.Fatalf("Could not connect to database %v", err)
	}

	e := echo.New()

	tasksRepo := taskservice.NewTaskRepository(database)
	tasksService := taskservice.NewTaskService(tasksRepo)
	tasksHandler := handlers.NewTaskHandler(tasksService)

	usersRepo := userservice.NewUserRepository(database)
	usersService := userservice.NewUserService(usersRepo)
	usersHandler := handlers.NewUserHandler(usersService)

	tasksStrictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	usersStrictHandler := users.NewStrictHandler(usersHandler, nil)

	tasks.RegisterHandlers(e, tasksStrictHandler)
	users.RegisterHandlers(e, usersStrictHandler)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	log.Fatal(e.Start(":8080"))
}

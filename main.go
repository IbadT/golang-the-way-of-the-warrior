package main

import (
	"net/http"

	"github.com/IbadT/golang-the-way-of-the-warrior.git/handlers"
)

func main() {
	http.HandleFunc("/tasks", handlers.GetTasks)
	http.HandleFunc("/task-id", handlers.GetTaskById)

	http.HandleFunc("/create/task", handlers.CreateTask)

	http.HandleFunc("/update/task", handlers.UpdateTaskById)
	http.HandleFunc("/update/task/completed", handlers.UpdateTaskCompletedById)

	http.HandleFunc("/delete/task-id", handlers.DeleteTaskById)

	// http.HandleFunc("/", handlers.GetUsers)
	// http.HandleFunc("/users", handlers.GetUsersByName)
	// http.HandleFunc("/:id", handlers.GetUserById)
	// http.HandleFunc("/user", handlers.CreateUser)
	// http.HandleFunc("/patch/user/:id", handlers.UpdateUserById)
	// http.HandleFunc("/delete/user/:id", handlers.DeleteUserById)

	http.ListenAndServe("localhost:8080", nil)
}

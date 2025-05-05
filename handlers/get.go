package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

var (
	users []User
	tasks []Task
	id    = 1
)

type Task struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Get tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	if err := HandleErrorMethod(r.Method, http.MethodGet); err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}
	if len(tasks) == 0 {
		// HandleResponse(w, nil, &Task{})
		HandleResponse(w, &tasks, nil)
		return
	}
	HandleResponse(w, &tasks, nil)
}

// Get task by id
func GetTaskById(w http.ResponseWriter, r *http.Request) {
	if err := HandleErrorMethod(r.Method, http.MethodGet); err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Неверный UUID", http.StatusBadRequest)
		return
	}
	for _, task := range tasks {
		if task.ID == id {
			HandleResponse(w, nil, &task)
			return
		}
	}
	// http.Error(w, fmt.Sprintf("Такого id: %s не найдено", id), http.StatusNoContent)
	http.Error(w, fmt.Sprintf("Задача с id: %s не найдена", id), http.StatusNotFound)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if err := HandleErrorMethod(r.Method, http.MethodGet); err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	if len(users) == 0 {
		json.NewEncoder(w).Encode("Список пользователей пуст")
	} else {
		json.NewEncoder(w).Encode(users)
		return
	}
}

// get user by id
func GetUserById(w http.ResponseWriter, r *http.Request) {
	if err := HandleErrorMethod(r.Method, http.MethodGet); err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	idByParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idByParam)
	if err != nil {
		http.Error(w, "Неверный id", http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.ID == id {
			json.NewEncoder(w).Encode(map[string]User{
				"user": user,
			})
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Пользователь с id: %d не найден", id),
	})
}

// get users by name
func GetUsersByName(w http.ResponseWriter, r *http.Request) {
	if err := HandleErrorMethod(r.Method, http.MethodGet); err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	name := strings.TrimSpace(r.URL.Query().Get("name"))

	for _, user := range users {
		if strings.EqualFold(name, user.Name) {
			fmt.Println(user)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]User{
				"user": user,
			})
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Пользователей с таким именем: %s нет", name),
	})
}

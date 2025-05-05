package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

// Update task by id
func UpdateTaskById(w http.ResponseWriter, r *http.Request) {
	if err := HandleErrorMethod(r.Method, http.MethodPut); err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Некорректный id", http.StatusBadRequest)
		return
	}

	var taskPayload struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&taskPayload); err != nil {
		http.Error(w, "Ошибка при декодировании JSON", http.StatusBadRequest)
		return
	}

	for index, task := range tasks {
		if task.ID == id {
			tasks[index] = Task{
				ID:          task.ID,
				Title:       taskPayload.Title,
				Description: taskPayload.Description,
				IsCompleted: task.IsCompleted,
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": fmt.Sprintf("Задача с id: %s обновлена", id),
			})
			return
		}
	}
	http.Error(w, fmt.Sprintf("Задача с id: %s не найдена", id), http.StatusNotFound)
}

// Update task (is_completed) by id
func UpdateTaskCompletedById(w http.ResponseWriter, r *http.Request) {
	if err := HandleErrorMethod(r.Method, http.MethodPatch); err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Некорректный id", http.StatusBadRequest)
		return
	}

	// var payload struct {
	// 	IsCompleted bool `json:"is_completed"`
	// }
	// if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
	// 	http.Error(w, "Ошибка при декодировании JSON", http.StatusBadRequest)
	// 	return
	// }

	for index, task := range tasks {
		if task.ID == id {
			tasks[index].IsCompleted = !task.IsCompleted
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": fmt.Sprintf("Задача с id: %s обновилась", id),
			})
			return
		}
	}

	http.Error(w, fmt.Sprintf("Задача с id: %s не найдена", id), http.StatusNotFound)
}

// Update user by id
func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	if err := HandleErrorMethod(r.Method, http.MethodPatch); err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	userIdStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Некорректный id", http.StatusBadRequest)
		return
	}

	var bodyUser User
	if err := json.NewDecoder(r.Body).Decode(&bodyUser); err != nil {
		http.Error(w, "Ошибка при декодировании JSON", http.StatusBadRequest)
		return
	}

	for index, user := range users {
		if user.ID == id {
			users[index].Name = bodyUser.Name
			json.NewEncoder(w).Encode(fmt.Sprintf("Пользователь с id %d обновлен", id))
			return
		}
	}

	http.Error(w, fmt.Sprintf("Пользователь с таким id %d не найден", id), http.StatusNotFound)
}

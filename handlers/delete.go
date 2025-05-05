package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

// Delete task by id
func DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	if err := HandleErrorMethod(r.Method, http.MethodDelete); err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Неверный UUID", http.StatusBadRequest)
		return
	}

	for index, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:index], tasks[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(fmt.Sprintf("Таска с id: %s удалена", id))
			return
		}
	}

	http.Error(w, fmt.Sprintf("Таска с id %d не найден", id), http.StatusNotFound)
}

// Delete user by id
func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	if err := HandleErrorMethod(r.Method, http.MethodDelete); err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	userIdParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(userIdParam)
	if err != nil {
		http.Error(w, "Неверный id", http.StatusBadRequest)
	}

	for index, user := range users {
		if user.ID == id {
			users = append(users[:index], users[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(fmt.Sprintf("Пользователь с id %d удален", id))
			return
		}
	}

	http.Error(w, fmt.Sprintf("Пользователь с id %d не найден", id), http.StatusNotFound)
}

// delete users by name
func DeleteUsersByName(w http.ResponseWriter, r *http.Request) {

}

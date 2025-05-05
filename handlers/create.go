package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

// Create task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	if err := HandleErrorMethod(r.Method, http.MethodPost); err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Ошибка при декодировании JSON", http.StatusBadRequest)
		return
	}

	if task.ID == uuid.Nil {
		task.ID = uuid.New()
	}
	tasks = append(tasks, task)

	w.WriteHeader(http.StatusCreated)
	HandleResponse(w, nil, &task)
}

// Create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if err := HandleErrorMethod(r.Method, http.MethodPost); err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Ошибка при декодировании JSON", http.StatusBadRequest)
		return
	}

	newUser.ID = id
	id++
	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

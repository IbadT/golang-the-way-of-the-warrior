package handlers

import (
	"encoding/json"
	"net/http"
)

func HandleResponse(w http.ResponseWriter, tasks *[]Task, task *Task) {
	w.Header().Set("Content-Type", "application/json")

	if tasks != nil && len(*tasks) == 0 {
		// json.NewEncoder(w).Encode(map[string]interface{}{
		// 	"tasks": []Task{},
		// })
		json.NewEncoder(w).Encode([]Task{})
		return
	}

	if task != nil {
		// json.NewEncoder(w).Encode(map[string]Task{
		// 	"task": *task,
		// })
		json.NewEncoder(w).Encode(*task)
		return
	}

	if tasks != nil {
		// json.NewEncoder(w).Encode(map[string]interface{}{
		// 	"tasks": *tasks,
		// })
		json.NewEncoder(w).Encode(*tasks)
		return
	}

	http.Error(w, "Некорректный запрос: отсутствуют данные", http.StatusBadRequest)
}

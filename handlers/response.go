package handlers

import (
	"encoding/json"
	"net/http"
)

func HandleResponse(w http.ResponseWriter, tasks *[]Task, task *Task) {
	w.Header().Set("Content-Type", "application/json")

	// ✅ Проверяем `tasks` перед использованием
	if tasks != nil && len(*tasks) == 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"tasks": []Task{},
		})
		return
	}

	// ✅ Проверяем `task` перед использованием
	if task != nil {
		json.NewEncoder(w).Encode(map[string]Task{
			"task": *task,
			// "task": []Task{},
		})
		return
	}

	// ✅ Проверяем, что `tasks` не `nil`
	if tasks != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"tasks": *tasks,
		})
		return
	}

	// Если `tasks` и `task` оба `nil`
	http.Error(w, "Некорректный запрос: отсутствуют данные", http.StatusBadRequest)
}

// func HandleResponse(w http.ResponseWriter, tasks *[]Task, task *Task) {
// 	w.Header().Set("Content-Type", "application/json")
// 	if len(*tasks) == 0 {
// 		json.NewEncoder(w).Encode(map[string]interface{}{
// 			"tasks": []Task{},
// 		})
// 		return
// 	}

// 	if task != nil {
// 		json.NewEncoder(w).Encode(map[string]Task{
// 			"task": *task,
// 		})
// 		return
// 	}

// 	if tasks == nil || len(*tasks) == 0 {
// 		json.NewEncoder(w).Encode(map[string]interface{}{
// 			"tasks": []Task{},
// 		})
// 	}

// 	json.NewEncoder(w).Encode(map[string][]Task{
// 		"tasks": *tasks,
// 	})

// }

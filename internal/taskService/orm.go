package taskservice

import "github.com/google/uuid"

type Task struct {
	ID     uuid.UUID `json:"id" gorm:"primaryKey"`
	Title  string    `json:"title"`
	IsDone bool      `json:"is_done"`
}

type UpdateTitleTaskRequest struct {
	Title string `json:"title"`
}

type RequestTaskBody struct {
	Title string `json:"title"`
}

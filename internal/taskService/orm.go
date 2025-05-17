package taskservice

import "github.com/google/uuid"

type Task struct {
	ID     uuid.UUID `json:"id" gorm:"primaryKey"`
	Title  string    `json:"title"`
	IsDone bool      `json:"is_done"`
	UserID uuid.UUID `json:"user_id"`
}

type UpdateTitleTaskRequest struct {
	Title string `json:"title"`
}

type RequestTaskBody struct {
	Title  string    `json:"title"`
	UserID uuid.UUID `json:"user_id"`
}

package taskservice

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Task struct {
	ID     uuid.UUID `gorm:"primaryKey" json:"id"`
	Title  string    `json:"title"`
	IsDone bool      `json:"is_done"`
}

type RequestBody struct {
	Title string `json:"title" validate:"required,min=3,max=100"`
}

var validate = validator.New()

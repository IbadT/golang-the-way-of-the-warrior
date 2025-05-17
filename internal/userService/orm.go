package userservice

import (
	taskservice "github.com/IbadT/golang-the-way-of-the-warrior.git/internal/taskService"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey"`
	Email    string
	Password string
	Tasks    []taskservice.Task
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserPasswordRequest struct {
	Password string `json:"password"`
}

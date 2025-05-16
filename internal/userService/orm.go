package userservice

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey"`
	Email    string
	Password string
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserPasswordRequest struct {
	Password string `json:"password"`
}

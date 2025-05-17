package userservice

import (
	taskservice "github.com/IbadT/golang-the-way-of-the-warrior.git/internal/taskService"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) error

	GetUsers() ([]User, error)
	GetUserByID(id uuid.UUID) (User, error)
	GetUserByEmail(email string) (User, error)
	GetTasksByUserID(user_id uuid.UUID) ([]taskservice.Task, error)

	UpdateUser(user User) error

	DeleteUserByID(id uuid.UUID) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) GetUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetTasksByUserID(user_id uuid.UUID) ([]taskservice.Task, error) {
	var tasks []taskservice.Task
	err := r.db.Find(&tasks, "user_id = ?", user_id).Error
	return tasks, err
}

func (r *userRepository) GetUserByID(id uuid.UUID) (User, error) {
	var user User
	err := r.db.First(&user).Error
	return user, err
}

func (r *userRepository) GetUserByEmail(email string) (User, error) {
	var user User
	err := r.db.First(&user, "email = ?", email).Error
	return user, err
}

func (r *userRepository) UpdateUser(user User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) DeleteUserByID(id uuid.UUID) error {
	return r.db.Delete(&User{}, id).Error
}

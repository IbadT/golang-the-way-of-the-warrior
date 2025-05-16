package userservice

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) error

	GetUsers() ([]User, error)
	GetUserByID(id uuid.UUID) (User, error)
	GetUserByEmail(email string) (User, error)

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

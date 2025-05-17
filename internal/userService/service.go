package userservice

import (
	"errors"
	"fmt"
	"sync"
	"time"

	taskservice "github.com/IbadT/golang-the-way-of-the-warrior.git/internal/taskService"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user UserRequest) (User, error)

	GetUsers() ([]User, error)
	GetUserByID(id uuid.UUID) (User, error)
	GetTasksByUserID(user_id uuid.UUID) ([]taskservice.Task, error)
	IncrementRequestsCount()

	UpdateUser(id uuid.UUID, body UserRequest) (User, error)
	UpdateUserPassword(id uuid.UUID, body UpdateUserPasswordRequest) (User, error)

	DeleteUserByID(id uuid.UUID) error
}

type userService struct {
	repo            UserRepository
	counterRequests uint16
	mu              sync.Mutex
}

func NewUserService(r UserRepository) *userService {
	return &userService{repo: r}
}

func (s *userService) IncrementRequestsCount() {
	s.mu.Lock()
	s.counterRequests++
	count := s.counterRequests
	s.mu.Unlock()
	fmt.Printf("\x1b[34m[%s] Запросов обработано: %d\x1b[0m\n", time.Now().Format("15:04:05"), count)
}

func (s *userService) CreateUser(body UserRequest) (User, error) {
	s.IncrementRequestsCount()

	existingUser, err := s.repo.GetUserByEmail(body.Email)
	if err == nil && existingUser.ID != uuid.Nil {
		return User{}, errors.New("пользователь с таким email уже существует")
	}

	user := User{
		ID:       uuid.New(),
		Email:    body.Email,
		Password: body.Password,
	}

	if err := s.repo.CreateUser(user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *userService) GetTasksByUserID(user_id uuid.UUID) ([]taskservice.Task, error) {
	return s.repo.GetTasksByUserID(user_id)
}

func (s *userService) GetUsers() ([]User, error) {
	s.IncrementRequestsCount()

	return s.repo.GetUsers()
}

func (s *userService) GetUserByID(id uuid.UUID) (User, error) {
	s.IncrementRequestsCount()

	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(id uuid.UUID, body UserRequest) (User, error) {
	s.IncrementRequestsCount()

	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return User{}, err
	}

	user.Email = body.Email
	user.Password = body.Password
	if err := s.repo.UpdateUser(user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *userService) UpdateUserPassword(id uuid.UUID, body UpdateUserPasswordRequest) (User, error) {
	s.IncrementRequestsCount()

	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return User{}, err
	}

	user.Password = body.Password
	if err := s.repo.UpdateUser(user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *userService) DeleteUserByID(id uuid.UUID) error {
	s.IncrementRequestsCount()

	return s.repo.DeleteUserByID(id)
}

package userservice

import (
	"errors"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user UserRequest) (User, error)
	GetUsers() ([]User, error)
	GetUserByID(id uuid.UUID) (User, error)
	UpdateUser(id uuid.UUID, body UserRequest) (User, error)
	UpdateUserPassword(id uuid.UUID, body UpdateUserPasswordRequest) (User, error)
	DeleteUserByID(id uuid.UUID) error
}

// добавить mutex

type userService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(body UserRequest) (User, error) {
	// првоеряем, зарегистрировал ли пользователь с таким email	!!!!!!!!
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

func (s *userService) GetUsers() ([]User, error) {
	return s.repo.GetUsers()
}

func (s *userService) GetUserByID(id uuid.UUID) (User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(id uuid.UUID, body UserRequest) (User, error) {
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
	return s.repo.DeleteUserByID(id)
}

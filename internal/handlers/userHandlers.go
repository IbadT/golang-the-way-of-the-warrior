package handlers

import (
	"context"

	userservice "github.com/IbadT/golang-the-way-of-the-warrior.git/internal/userService"
	"github.com/IbadT/golang-the-way-of-the-warrior.git/internal/web/users"
)

type UserHandler struct {
	service userservice.UserService
}

func NewUserHandler(s userservice.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// CreateUser implements users.StrictServerInterface.
func (u *UserHandler) CreateUser(ctx context.Context, request users.CreateUserRequestObject) (users.CreateUserResponseObject, error) {
	body := request.Body
	userToCreate := userservice.UserRequest{
		Email:    *body.Email,
		Password: *body.Password,
	}

	createdUser, err := u.service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}

	response := users.CreateUser201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	return response, nil
}

// DeleteUserByID implements users.StrictServerInterface.
func (u *UserHandler) DeleteUserByID(ctx context.Context, request users.DeleteUserByIDRequestObject) (users.DeleteUserByIDResponseObject, error) {
	id := request.Id
	if err := u.service.DeleteUserByID(id); err != nil {
		return nil, err
	}
	return users.DeleteUserByID204Response{}, nil
}

// GetUserByID implements users.StrictServerInterface.
func (u *UserHandler) GetUserByID(ctx context.Context, request users.GetUserByIDRequestObject) (users.GetUserByIDResponseObject, error) {
	id := request.Id
	user, err := u.service.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	response := users.GetUserByID200JSONResponse{
		Id:       &user.ID,
		Email:    &user.Email,
		Password: &user.Password,
	}
	return response, nil
}

// GetUsers implements users.StrictServerInterface.
func (u *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	gettedUsers, err := u.service.GetUsers()
	if err != nil {
		return nil, err
	}

	// response := users.GetTasks200JSONResponse{}
	response := users.GetUsers200JSONResponse{}
	for _, usr := range gettedUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

// UpdateUser implements users.StrictServerInterface.
func (u *UserHandler) UpdateUser(ctx context.Context, request users.UpdateUserRequestObject) (users.UpdateUserResponseObject, error) {
	id := request.Id
	body := request.Body
	userRequest := userservice.UserRequest{
		Email:    *body.Email,
		Password: *body.Password,
	}

	user, err := u.service.UpdateUser(id, userRequest)
	if err != nil {
		return nil, err
	}

	response := users.UpdateUser200JSONResponse{
		Email:    &user.Email,
		Password: &user.Password,
	}
	return response, nil
}

// UpdateUserPassword implements users.StrictServerInterface.
func (u *UserHandler) UpdateUserPassword(ctx context.Context, request users.UpdateUserPasswordRequestObject) (users.UpdateUserPasswordResponseObject, error) {
	id := request.Id
	body := request.Body
	updateUserPasswordData := userservice.UpdateUserPasswordRequest{
		Password: *body.Password,
	}

	user, err := u.service.UpdateUserPassword(id, updateUserPasswordData)
	if err != nil {
		return nil, err
	}

	response := users.UpdateUserPassword200JSONResponse{
		Id:       &user.ID,
		Email:    &user.Email,
		Password: &user.Password,
	}

	return response, nil
}

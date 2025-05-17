package handlers

import (
	"context"

	taskservice "github.com/IbadT/golang-the-way-of-the-warrior.git/internal/taskService"
	"github.com/IbadT/golang-the-way-of-the-warrior.git/internal/web/tasks"
)

type TaskHandler struct {
	service taskservice.TaskService
}

func NewTaskHandler(s taskservice.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

// CreateTask implements tasks.StrictServerInterface.
func (t *TaskHandler) CreateTask(ctx context.Context, request tasks.CreateTaskRequestObject) (tasks.CreateTaskResponseObject, error) {
	body := request.Body

	taskToCreate := taskservice.RequestTaskBody{
		UserID: *body.UserId,
		Title:  *body.Title,
	}

	task, err := t.service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.CreateTask201JSONResponse{
		Id:     &task.ID,
		Title:  &task.Title,
		IsDone: &task.IsDone,
		UserId: &task.UserID,
	}

	return response, nil
}

// GetTasksByUserID implements tasks.StrictServerInterface.
func (t *TaskHandler) GetTasksByUserID(ctx context.Context, request tasks.GetTasksByUserIDRequestObject) (tasks.GetTasksByUserIDResponseObject, error) {
	userId := request.UserId

	allTasks, err := t.service.GetTasksByUserID(userId)
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasksByUserID200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Title:  &tsk.Title,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}

	return response, nil
}

// DeleteTaskById implements tasks.StrictServerInterface.
func (t *TaskHandler) DeleteTaskById(ctx context.Context, request tasks.DeleteTaskByIdRequestObject) (tasks.DeleteTaskByIdResponseObject, error) {
	id := request.Id

	err := t.service.DeleteTaskById(id)
	if err != nil {
		return nil, err
	}

	return tasks.DeleteTaskById204Response{}, nil
}

// GetTaskById implements tasks.StrictServerInterface.
func (t *TaskHandler) GetTaskById(ctx context.Context, request tasks.GetTaskByIdRequestObject) (tasks.GetTaskByIdResponseObject, error) {
	id := request.Id

	task, err := t.service.GetTaskById(id)
	if err != nil {
		return nil, err
	}

	response := tasks.GetTaskById200JSONResponse{
		Id:     &task.ID,
		Title:  &task.Title,
		IsDone: &task.IsDone,
		UserId: &task.UserID,
	}

	return response, nil
}

// GetTasks implements tasks.StrictServerInterface.
func (t *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	var isDone *bool
	if request.Params.IsDone != nil {
		isDone = request.Params.IsDone
	}

	allTasks, err := t.service.GetTasks(isDone)
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Title:  &tsk.Title,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}

	return response, nil
}

// UpdateTaskCompletedById implements tasks.StrictServerInterface.
func (t *TaskHandler) UpdateTaskCompletedById(ctx context.Context, request tasks.UpdateTaskCompletedByIdRequestObject) (tasks.UpdateTaskCompletedByIdResponseObject, error) {
	id := request.Id

	task, err := t.service.UpdateTaskCompletedById(id)
	if err != nil {
		return nil, err
	}

	response := tasks.UpdateTaskCompletedById200JSONResponse{
		Id:     &task.ID,
		Title:  &task.Title,
		IsDone: &task.IsDone,
	}

	return response, nil
}

// UpdateTitleTaskById implements tasks.StrictServerInterface.
func (t *TaskHandler) UpdateTitleTaskById(ctx context.Context, request tasks.UpdateTitleTaskByIdRequestObject) (tasks.UpdateTitleTaskByIdResponseObject, error) {
	id := request.Id

	body := taskservice.UpdateTitleTaskRequest{
		Title: *request.Body.Title,
	}

	task, err := t.service.UpdateTitleTaskById(id, body)
	if err != nil {
		return nil, err
	}

	response := tasks.UpdateTitleTaskById200JSONResponse{
		Id:     &task.ID,
		Title:  &task.Title,
		IsDone: &task.IsDone,
	}

	return response, nil
}

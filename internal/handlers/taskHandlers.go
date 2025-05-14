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

// GetTasks implements tasks.StrictServerInterface.
func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	var isDone *bool
	if request.Params.IsDone != nil {
		isDone = request.Params.IsDone
	}
	allTasks, err := h.service.GetTasks(isDone)
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Title:  &tsk.Title,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

// GetTaskByID implements tasks.StrictServerInterface.
func (h *TaskHandler) GetTaskByID(ctx context.Context, request tasks.GetTaskByIDRequestObject) (tasks.GetTaskByIDResponseObject, error) {
	idParam := request.Id
	task, err := h.service.GetTaskById(idParam)
	if err != nil {
		return nil, err
	}
	response := tasks.GetTaskByID200JSONResponse{
		Id:     &task.ID,
		Title:  &task.Title,
		IsDone: &task.IsDone,
	}
	return response, nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	body := request.Body

	taskToCreate := taskservice.RequestBody{
		Title: body.Title,
	}
	createdTask, err := h.service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Title:  &createdTask.Title,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}

// UpdateTaskCompletedById implements tasks.StrictServerInterface.
func (h *TaskHandler) UpdateTaskCompletedById(ctx context.Context, request tasks.UpdateTaskCompletedByIdRequestObject) (tasks.UpdateTaskCompletedByIdResponseObject, error) {
	id := request.Id
	task, err := h.service.UpdateTaskCompletedById(id)
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

// UpdateTaskByID implements tasks.StrictServerInterface.
func (h *TaskHandler) UpdateTaskByID(ctx context.Context, request tasks.UpdateTaskByIDRequestObject) (tasks.UpdateTaskByIDResponseObject, error) {
	id := request.Id
	body := request.Body
	taskToUpdateTitle := taskservice.RequestBody{
		Title: body.Title,
	}
	task, err := h.service.UpdateTitleTaskById(id, taskToUpdateTitle)
	if err != nil {
		return nil, err
	}
	response := tasks.UpdateTaskByID200JSONResponse{
		Id:     &task.ID,
		Title:  &task.Title,
		IsDone: &task.IsDone,
	}
	return response, nil
}

// DeleteTaskById implements tasks.StrictServerInterface.
func (h *TaskHandler) DeleteTaskById(ctx context.Context, request tasks.DeleteTaskByIdRequestObject) (tasks.DeleteTaskByIdResponseObject, error) {
	id := request.Id
	err := h.service.DeleteTaskById(id)
	if err != nil {
		return nil, err
	}

	return tasks.DeleteTaskById204Response{}, nil
}

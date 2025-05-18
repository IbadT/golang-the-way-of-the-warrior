package taskservice

import (
	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(body RequestTaskBody) (Task, error)
	GetTasks(isDone *bool) ([]Task, error)
	// GetTasksByUserID(user_id uuid.UUID) ([]Task, error)
	GetTaskById(id uuid.UUID) (Task, error)
	UpdateTaskCompletedById(id uuid.UUID, body RequestIsDoneBody) (Task, error)
	UpdateTitleTaskById(id uuid.UUID, body UpdateTitleTaskRequest) (Task, error)
	DeleteTaskById(id uuid.UUID) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(body RequestTaskBody) (Task, error) {
	task := Task{
		ID:     uuid.New(),
		Title:  body.Title,
		IsDone: false,
		UserID: body.UserID,
	}

	if err := s.repo.CreateTask(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *taskService) GetTasks(isDone *bool) ([]Task, error) {
	if isDone == nil {
		return s.repo.GetTasks()
	}
	return s.repo.GetTasksByCompleted(*isDone)
}

func (s *taskService) GetTaskById(id uuid.UUID) (Task, error) {
	return s.repo.GetTaskById(id)
}

func (s *taskService) UpdateTaskCompletedById(id uuid.UUID, body RequestIsDoneBody) (Task, error) {
	task, err := s.repo.GetTaskById(id)
	if err != nil {
		return Task{}, err
	}

	task.IsDone = body.IsDone
	if err := s.repo.UpdateTask(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *taskService) UpdateTitleTaskById(id uuid.UUID, body UpdateTitleTaskRequest) (Task, error) {
	task, err := s.repo.GetTaskById(id)

	if err != nil {
		return Task{}, err
	}
	task.Title = body.Title
	if err := s.repo.UpdateTask(task); err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *taskService) DeleteTaskById(id uuid.UUID) error {
	return s.repo.DeleteTaskById(id)
}

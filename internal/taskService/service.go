package taskservice

import (
	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(body RequestBody) (Task, error)
	GetTasks(isDoneQuery *bool) ([]Task, error)
	GetTaskById(id uuid.UUID) (Task, error)
	UpdateTaskCompletedById(id uuid.UUID) (Task, error)
	UpdateTitleTaskById(id uuid.UUID, body RequestBody) (Task, error)
	DeleteTaskById(id uuid.UUID) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(body RequestBody) (Task, error) {
	if err := validate.Struct(body); err != nil {
		return Task{}, err
	}

	task := Task{
		ID:     uuid.New(),
		Title:  body.Title,
		IsDone: false,
	}

	if err := s.repo.CreateTask(task); err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *taskService) GetTasks(isDoneQuery *bool) ([]Task, error) {
	if isDoneQuery == nil {
		return s.repo.GetTasks()
	}
	return s.repo.GetTasksByCompleted(*isDoneQuery)
}

func (s *taskService) GetTaskById(id uuid.UUID) (Task, error) {
	return s.repo.GetTaskById(id)
}

func (s *taskService) UpdateTaskCompletedById(id uuid.UUID) (Task, error) {
	task, err := s.repo.GetTaskById(id)
	if err != nil {
		return Task{}, err
	}

	task.IsDone = !task.IsDone
	if err := s.repo.UpdateTaskCompletedById(id, task); err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *taskService) UpdateTitleTaskById(id uuid.UUID, body RequestBody) (Task, error) {
	if err := validate.Struct(body); err != nil {
		return Task{}, err
	}

	task, err := s.repo.GetTaskById(id)
	if err != nil {
		return Task{}, err
	}

	task.Title = body.Title
	if err := s.repo.UpdateTitleTaskById(id, task); err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *taskService) DeleteTaskById(id uuid.UUID) error {
	return s.repo.DeleteTaskById(id)
}

package taskservice

import (
	"strconv"

	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(body RequestBody) (Task, error)
	GetTasks(isDoneQuery string) ([]Task, error)
	GetTaskById(idStr string) (Task, error)
	UpdateTaskCompletedById(idStr string) (Task, error)
	UpdateTitleTaskById(idStr string, body RequestBody) (Task, error)
	DeleteTaskById(idStr string) error
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

func (s *taskService) GetTasks(isDoneQuery string) ([]Task, error) {
	if isDoneQuery != "" {
		isDone, err := strconv.ParseBool(isDoneQuery)
		if err != nil {
			return []Task{}, err
		}
		return s.repo.GetTasksByCompleted(isDone)
	}
	return s.repo.GetTasks()
}

func (s *taskService) GetTaskById(idStr string) (Task, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return Task{}, err
	}
	return s.repo.GetTaskById(id)
}

func (s *taskService) UpdateTaskCompletedById(idStr string) (Task, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return Task{}, err
	}

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

func (s *taskService) UpdateTitleTaskById(idStr string, body RequestBody) (Task, error) {
	if err := validate.Struct(body); err != nil {
		return Task{}, err
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
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

func (s *taskService) DeleteTaskById(idStr string) error {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return err
	}
	return s.repo.DeleteTaskById(id)
}

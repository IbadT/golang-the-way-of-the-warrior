package taskservice

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task Task) error

	GetTasks() ([]Task, error)
	GetTasksByCompleted(isDone bool) ([]Task, error)
	GetTaskById(id uuid.UUID) (Task, error)
	GetTasksByUserID(user_id uuid.UUID) ([]Task, error)

	UpdateTask(task Task) error

	DeleteTaskById(id uuid.UUID) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) error {
	return r.db.Create(&task).Error
}

func (r *taskRepository) GetTasksByUserID(user_id uuid.UUID) ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks, "user_id = ?", user_id).Error
	return tasks, err
}

func (r *taskRepository) GetTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTasksByCompleted(isDone bool) ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks, "is_done = ?", isDone).Error

	return tasks, err
}

func (r *taskRepository) GetTaskById(id uuid.UUID) (Task, error) {
	var task Task
	err := r.db.First(&task, "id = ?", id).Error
	return task, err
}

func (r *taskRepository) UpdateTask(task Task) error {
	return r.db.Save(&task).Error
}

func (r *taskRepository) DeleteTaskById(id uuid.UUID) error {
	return r.db.Delete(&Task{}, id).Error
}

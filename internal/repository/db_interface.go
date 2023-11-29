// internal/repository/db_interface.go

package repository

import "backend-go/internal/model"

// TaskRepository is an interface for database operations related to tasks.
type TaskRepository interface {
	ConnectDB() error
	GetAllTasks() ([]model.Task, error)
	GetTaskByID(id int) (*model.Task, error)
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task) error
	DeleteTask(id int) error
}

// internal/repository/db_interface.go

package repository

// Database is an interface for database operations related to tasks.
type Database interface {
	ConnectDB() error
	// GetAllTasks() ([]model.Task, error)
	// GetTaskByID(id int) (*model.Task, error)
	// CreateTask(task *model.Task) error
	// UpdateTask(task *model.Task) error
	// DeleteTask(id int) error
}

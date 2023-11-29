// internal/model/model.go

package model

// Task represents a task in the task management system.
type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	IsFav     bool   `json:"isFav"`
	Completed bool   `json:"completed"`
	Notes     string `json:"notes"`
}

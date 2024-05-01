// internal/router/router.go

package router

import (
	"backend-server/internal/repository"

	"github.com/gorilla/mux"
)

// NewRouter creates a new instance of the HTTP router with configured routes.
func NewRouter(taskRepo repository.Database) *mux.Router {
	r := mux.NewRouter()

	// Endpoints
	// r.HandleFunc("/tasks", handlers.GetTasksHandler(taskRepo)).Methods("GET")
	// r.HandleFunc("/task/{id}", handlers.GetTaskHandler(taskRepo)).Methods("GET")
	// r.HandleFunc("/task", handlers.CreateTaskHandler(taskRepo)).Methods("POST")
	// r.HandleFunc("/task/{id}", handlers.UpdateTaskHandler(taskRepo)).Methods("PUT")
	// r.HandleFunc("/task/{id}", handlers.PatchTaskHandler(taskRepo)).Methods("PATCH")
	// r.HandleFunc("/task/{id}", handlers.DeleteTaskHandler(taskRepo)).Methods("DELETE")
	// r.HandleFunc("/tasks", handlers.DeleteTasksHandler(taskRepo)).Methods("DELETE")

	return r
}

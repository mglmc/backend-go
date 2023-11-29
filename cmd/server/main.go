package main

import (
	"backend-server/internal/model"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"

	"backend-go/internal/model"
)

// TaskList structure to hold a list of tasks
type TaskList struct {
	Tasks []model.Task `json:"tasks"`
	mu    sync.Mutex
}

var taskList TaskList

func init() {
	// Load tasks from JSON file on startup
	loadTasks()
}

func main() {
	r := mux.NewRouter()

	// Endpoints
	r.HandleFunc("/tasks", getTasksHandler).Methods("GET")
	r.HandleFunc("/task/{id}", getTaskHandler).Methods("GET")
	r.HandleFunc("/task", createTaskHandler).Methods("POST")
	r.HandleFunc("/task/{id}", updateTaskHandler).Methods("PUT")
	r.HandleFunc("/task/{id}", patchTaskHandler).Methods("PATCH")
	r.HandleFunc("/task/{id}", deleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/tasks", deleteTasksHandler).Methods("DELETE")

	// Start the server
	serverAddr := ":8080"
	srv := &http.Server{
		Handler:      r,
		Addr:         serverAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Server listening on %s\n", serverAddr)
	log.Fatal(srv.ListenAndServe())
}

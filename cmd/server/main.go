package main

import (
	"backend-server/internal/handler"

	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Endpoints
	r.HandleFunc("/tasks", handler.GetTasksHandler).Methods("GET")
	r.HandleFunc("/task/{id}", handler.GetTaskHandler).Methods("GET")
	r.HandleFunc("/task", handler.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/task/{id}", handler.UpdateTaskHandler).Methods("PUT")
	r.HandleFunc("/task/{id}", handler.PatchTaskHandler).Methods("PATCH")
	r.HandleFunc("/task/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/tasks", handler.DeleteTasksHandler).Methods("DELETE")

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

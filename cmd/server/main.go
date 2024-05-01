package main

import (
	"backend-server/internal/handler"
	"context"
	"os"
	"os/signal"
	"syscall"

	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	serverAddress    = ":8080"
	webServerAddress = ":80"
	serverTimeout    = 15 * time.Second
)

func main() {
	webDir := "../../static"
	landingWebDir := webDir + "/landing/dist"
	delaporteWebDir := webDir + "/delaporte"

	rServer := registerRoutes()

	rWeb := mux.NewRouter()
	info, err := os.Stat(delaporteWebDir + "/index.html")
	if err != nil {
		log.Printf("Cannot access index.html in %s: %v", delaporteWebDir, err)
	} else {
		log.Printf("Permissions of index.html in %s: %v", delaporteWebDir, info.Mode())
	}
	rWeb.PathPrefix("/delaporte").Handler(http.StripPrefix("/delaporte", http.FileServer(http.Dir(delaporteWebDir))))
	log.Printf("Serving Delaporte from %s", delaporteWebDir)
	rWeb.PathPrefix("/").Handler(http.FileServer(http.Dir(landingWebDir)))
	log.Printf("Serving landing page from %s", landingWebDir)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	srv, err := startServer(serverAddress, rServer)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	webServer, err := startServer(webServerAddress, rWeb)
	if err != nil {
		log.Fatalf("Failed to start web server: %v", err)
	}

	<-stop
	fmt.Println("Shutting down servers...")

	ctx, cancel := context.WithTimeout(context.Background(), serverTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown failed: %v", err)
	}

	if err := webServer.Shutdown(ctx); err != nil {
		log.Printf("Web server shutdown failed: %v", err)
	}

	log.Println("Servers gracefully stopped")
}

func startServer(addr string, handler http.Handler) (*http.Server, error) {
	srv := &http.Server{
		Handler:      handler,
		Addr:         addr,
		WriteTimeout: serverTimeout,
		ReadTimeout:  serverTimeout,
	}

	go func() {
		log.Printf("Starting server on %s", addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Error starting server on %s: %v", addr, err)
		}
	}()

	return srv, nil
}

func registerRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", handler.GetTasksHandler).Methods("GET")
	r.HandleFunc("/task/{id}", handler.GetTaskHandler).Methods("GET")
	r.HandleFunc("/task", handler.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/task/{id}", handler.UpdateTaskHandler).Methods("PUT")
	r.HandleFunc("/task/{id}", handler.PatchTaskHandler).Methods("PATCH")
	r.HandleFunc("/task/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/tasks", handler.DeleteTasksHandler).Methods("DELETE")

	return r
}

package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	taskList.mu.Lock()
	defer taskList.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskList)
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	taskList.mu.Lock()
	defer taskList.mu.Unlock()

	for _, task := range taskList.Tasks {
		if fmt.Sprint(task.ID) == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	http.NotFound(w, r)
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskList.mu.Lock()
	defer taskList.mu.Unlock()

	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Assign a unique ID and add the task
	newTask.ID = len(taskList.Tasks) + 1
	taskList.Tasks = append(taskList.Tasks, newTask)

	saveTasks()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	taskList.mu.Lock()
	defer taskList.mu.Unlock()

	for i, task := range taskList.Tasks {
		if fmt.Sprint(task.ID) == taskID {
			var updatedTask Task
			err := json.NewDecoder(r.Body).Decode(&updatedTask)
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			// Update the task
			taskList.Tasks[i] = updatedTask

			saveTasks()

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}

	http.NotFound(w, r)
}

func patchTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	taskList.mu.Lock()
	defer taskList.mu.Unlock()

	for i, task := range taskList.Tasks {
		if fmt.Sprint(task.ID) == taskID {
			var patchTask map[string]interface{}
			err := json.NewDecoder(r.Body).Decode(&patchTask)
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			// Apply partial updates
			if title, ok := patchTask["title"]; ok {
				taskList.Tasks[i].Title = title.(string)
			}
			if isFav, ok := patchTask["isFav"]; ok {
				taskList.Tasks[i].IsFav = isFav.(bool)
			}
			if completed, ok := patchTask["completed"]; ok {
				taskList.Tasks[i].Completed = completed.(bool)
			}
			if notes, ok := patchTask["notes"]; ok {
				taskList.Tasks[i].Notes = notes.(string)
			}

			saveTasks()

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(taskList.Tasks[i])
			return
		}
	}

	http.NotFound(w, r)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	taskList.mu.Lock()
	defer taskList.mu.Unlock()

	for i, task := range taskList.Tasks {
		if fmt.Sprint(task.ID) == taskID {
			// Remove the task
			taskList.Tasks = append(taskList.Tasks[:i], taskList.Tasks[i+1:]...)

			saveTasks()

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.NotFound(w, r)
}

func deleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	taskList.mu.Lock()
	defer taskList.mu.Unlock()

	// Clear all tasks
	taskList.Tasks = nil

	saveTasks()

	w.WriteHeader(http.StatusNoContent)
}

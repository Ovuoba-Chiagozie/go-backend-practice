package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks []Task
var mu sync.Mutex

var NextIDD = 1

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Welcome to my Task API")
		return
	}

	http.Error(w, "Incorrect method my boy", http.StatusMethodNotAllowed)
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		mu.Lock()
		defer mu.Unlock()

		var task Task
		err := json.NewDecoder(r.Body).Decode(&task)

		if err != nil {
			http.Error(w, "Invalid Json baby girl", http.StatusBadRequest)
			return
		}

		task.ID = NextIDD
		NextIDD++

		tasks = append(tasks, task)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(task)
		return
	}

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
		return
	}

	http.Error(w, "Method not allowed mate", http.StatusMethodNotAllowed)

}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Invalid ID Value", http.StatusBadRequest)
		return
	}

	for _, task := range tasks {
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	http.Error(w, "Task Not Found", http.StatusNotFound)
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Invalid ID value", http.StatusBadRequest)
		return
	}

	var taskBody Task
	err = json.NewDecoder(r.Body).Decode(&taskBody)

	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	for index, task := range tasks {
		if task.ID == id {
			pointer := &tasks
			(*pointer)[index] = taskBody
			(*pointer)[index].ID = id
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tasks[index])
			return
		}
	}

	http.Error(w, "Task NOt Found", http.StatusNotFound)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Invalid ID value", http.StatusBadRequest)
		return
	}

	var taskBody Task
	err = json.NewDecoder(r.Body).Decode(&taskBody)

	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	for index, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:index], tasks[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	http.Error(w, "Task NOt Found", http.StatusNotFound)
}

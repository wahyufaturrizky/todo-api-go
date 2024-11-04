package main

import (
    "net/http"
    "github.com/gorilla/mux"
		"encoding/json"
)

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todos = append(todos, todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func getTodoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range todos {
			if item.ID == params["id"] {
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(item)
					return
			}
	}
	w.WriteHeader(http.StatusNotFound)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range todos {
			if item.ID == params["id"] {
					todos = append(todos[:index], todos[index+1:]...)
					var todo Todo
					_ = json.NewDecoder(r.Body).Decode(&todo)
					todo.ID = params["id"]
					todos = append(todos, todo)
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(todo)
					return
			}
	}
	w.WriteHeader(http.StatusNotFound)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range todos {
			if item.ID == params["id"] {
					todos = append(todos[:index], todos[index+1:]...)
					break
			}
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos", createTodo).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
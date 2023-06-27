package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

var todos []Todo

func main() {
	// Initialize some sample TODO items
	todos = append(todos, Todo{ID: "1", Title: "Learn Go", IsDone: false})
	todos = append(todos, Todo{ID: "2", Title: "Build an API", IsDone: false})

	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Define API endpoints
	router.HandleFunc("/todos", GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", GetTodo).Methods("GET")
	router.HandleFunc("/todos", CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")

	// Start the HTTP server
	port := ":8080"
	log.Println("Server listening on http://localhost", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// GetTodos returns all TODO items
func GetTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

// GetTodo returns a specific TODO item by ID
func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for _, todo := range todos {
		if todo.ID == id {
			json.NewEncoder(w).Encode(todo)
			return
		}
	}

	http.NotFound(w, r)
}

// CreateTodo creates a new TODO item
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	todo.ID = "temp" // You can use a unique ID generation logic here

	todos = append(todos, todo)

	json.NewEncoder(w).Encode(todo)
}

// UpdateTodo updates a TODO item by ID
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for index, todo := range todos {
		if todo.ID == id {
			todos[index].IsDone = true
			json.NewEncoder(w).Encode(todos[index])
			return
		}
	}

	http.NotFound(w, r)
}

// DeleteTodo deletes a TODO item by ID
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for index, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.NotFound(w, r)
}

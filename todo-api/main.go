package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type TodoItem struct {
	ID          int
	Description string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type TodoList struct {
	items []*TodoItem
}

var templates = template.Must(template.ParseFiles("todos.html"))

func (tl *TodoList) load() {
	data, err := os.ReadFile("todos.json")
	if err != nil {
		log.Fatal("Error reading todos.json", err)
	}
	var todoList TodoList
	json.Unmarshal(data, &todoList.items)
	if err != nil {
		log.Fatal("Error reading json", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	todoList := &TodoList{}

	err := templates.ExecuteTemplate(w, "todos.html", todoList.items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
}

func update(w http.ResponseWriter, r *http.Request) {
}

func destroy(w http.ResponseWriter, r *http.Request) {
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		index(w, r)
	case "POST":
		create(w, r)
	case "PATCH":
		update(w, r)
	case "DELETE":
		destroy(w, r)
	}
}

func main() {
	http.HandleFunc("/todos/", resourceHandler)
	log.Println("Listening on http://localhost:3000/todos")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

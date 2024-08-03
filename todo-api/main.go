package main

import (
	"encoding/json"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"
)

type TodoItem struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Done        bool       `json:"done"`
	CreatedAt   time.Time  `json:"createdAt"`
	CompletedAt *time.Time `json:"completedAt"`
}

type TodoList struct {
	items []*TodoItem
}

var templates = template.Must(template.ParseFiles("todos.html", "create-todo.html"))

func (tl *TodoList) load() {
	data, err := os.ReadFile("todos.json")
	if err != nil {
		log.Fatal("Error reading todos.json", err)
	}

	if err := json.Unmarshal(data, &tl.items); err != nil {
		log.Fatal("Error unmarshaling", err)
	}
}

func (tl *TodoList) save(todo *TodoItem) {
	tl.items = append(tl.items, todo)
	json, err := json.Marshal(&tl.items)
	if err != nil {
		log.Fatal("save err while marshaling", err)
	}

	err = os.WriteFile("todos.json", json, fs.ModePerm)
	if err != nil {
		log.Fatal("save err while writing file", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	todoList := &TodoList{}
	todoList.load()
	err := templates.ExecuteTemplate(w, "todos.html", todoList.items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
	}

	todoList := &TodoList{}
	todoList.load()

	id := len(todoList.items) + 1
	todoItem := TodoItem{
		ID:          id,
		Description: r.FormValue("item"),
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	todoList.save(&todoItem)

	err = templates.ExecuteTemplate(w, "todos.html", todoList.items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

var templates = template.Must(template.ParseFiles("todos.html", "create-todo.html", "delete-todo.html"))

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

func destroy(w http.ResponseWriter, r *http.Request) {
	todoList := &TodoList{}
	todoList.load()
	id, err := strconv.Atoi(r.FormValue("todo_id"))
	if err != nil {
		log.Fatal("Problem converting string to integer", err)
	}
	todoList.complete(id)
	todoList.save(nil)

	err = templates.ExecuteTemplate(w, "todos.html", todoList.items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func update(w http.ResponseWriter, r *http.Request) {
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	if r.FormValue("_method") == "DELETE" {
		method = "DELETE"
	}

	switch method {
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

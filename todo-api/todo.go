package main

import (
	"encoding/json"
	"io/fs"
	"log"
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
	if todo != nil {
		tl.items = append(tl.items, todo)
	}

	json, err := json.Marshal(&tl.items)
	if err != nil {
		log.Fatal("save err while marshaling", err)
	}

	err = os.WriteFile("todos.json", json, fs.ModePerm)
	if err != nil {
		log.Fatal("save err while writing file", err)
	}
}

func (tl *TodoList) complete(id int) {
	for i := 0; i < len(tl.items); i++ {
		item := tl.items[i]
		if item.ID == id {
			now := time.Now()
			item.Done = true
			item.CompletedAt = &now
		}
	}

}

func (tl *TodoList) update(id int, desc string) {
	for i := 0; i < len(tl.items); i++ {
		item := tl.items[i]
		if item.ID == id {
			item.Description = desc
		}
	}
}

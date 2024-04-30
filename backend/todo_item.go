package main

type TodoItem struct {
	Title string `json:"title" db:"title"`
}

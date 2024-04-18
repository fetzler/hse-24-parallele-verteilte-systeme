package main

type TodoItem struct {
	Id    uint32 `json:"id" db:"task_id"`
	Title string `json:"title" db:"title"`
}

package main

type TodoItem struct {
<<<<<<< HEAD
	Id    uint32 `json:"id" db:"task_id"`
	Title string `json:"title" db:"title"`
=======
	Id    uint32 `json:"Id"`
	Title string `json:"Title"`
>>>>>>> add api endpoint to creat todo item
}

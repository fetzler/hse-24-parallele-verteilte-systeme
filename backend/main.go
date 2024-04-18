package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func createTodoItem(c *gin.Context) {
	_, err := conn.Query(context.Background(), "INSERT INTO tasks (title) VALUES ($1)", c.Param("title"))
	if err != nil {
		log.Println(err)
	}
	c.String(http.StatusOK, "TodoItem was added")
}

var conn *pgx.Conn

func main() {
	var err error
	conn, err = pgx.Connect(context.Background(), "postgres://postgres:postgres@db:5432/todo_list")
	if err != nil {
		fmt.Println(err)
		log.Fatalln("error connecting to the database.")
	}
	defer conn.Close(context.Background())
	router := gin.Default()
	router.POST("/todos/:title", createTodoItem)
	router.Run()
}

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
	rows, err := conn.Query(context.Background(), "INSERT INTO tasks (title) VALUES ($1)", c.Param("title"))
	rows.Close()
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	c.String(http.StatusCreated, http.StatusText(http.StatusCreated))
}

func updateTodoItem(c *gin.Context) {
	rows, err := conn.Query(context.Background(), "UPDATE tasks SET title = $1 WHERE task_id = $2", c.Param("title"), c.Param("id"))
	rows.Close()
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	c.String(http.StatusOK, http.StatusText(http.StatusOK))
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
	router.PUT("/todos/:id/:title", updateTodoItem)
	router.Run()
}

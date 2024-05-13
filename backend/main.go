package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func getAllTodoItems(c *gin.Context) {
	rows, _ := conn.Query(context.Background(), "SELECT title FROM tasks")
	defer rows.Close()

	var todos []string

	for rows.Next() {
		var todo string
		rows.Scan(&todo)

		todos = append(todos, todo)
	}

	c.JSON(http.StatusOK, todos)
}

func createTodoItem(c *gin.Context) {
	_, err := conn.Exec(context.Background(), "INSERT INTO tasks (title) VALUES ($1)", c.Param("title"))
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	c.String(http.StatusCreated, http.StatusText(http.StatusCreated))
}

func deleteTodoItem(c *gin.Context) {
	_, err := conn.Exec(context.Background(), "DELETE FROM tasks WHERE title = $1", c.Param("title"))
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, http.StatusText((http.StatusInternalServerError)))
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
	router.Use(cors.Default())

	router.GET("/todos", getAllTodoItems)
	router.POST("/todos/:title", createTodoItem)
	router.DELETE("/todos/:title", deleteTodoItem)
	router.Run()
}

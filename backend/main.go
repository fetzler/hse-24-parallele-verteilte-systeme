package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func createTodoItem(c *gin.Context) {
	_, err := conn.Exec(context.Background(), "INSERT INTO tasks (title) VALUES ($1)", c.Param("title"))
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	c.String(http.StatusCreated, http.StatusText(http.StatusCreated))
}

func updateTodoItem(c *gin.Context) {
	rows, err := conn.Exec(context.Background(), "UPDATE tasks SET title = $1 WHERE task_id = $2", c.Param("title"), c.Param("id"))
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	if rows.RowsAffected() == 0 {
		log.Println("No TodoItem found for this ID")
		c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
	c.String(http.StatusOK, http.StatusText(http.StatusOK))
}

func getTodoItemById(c *gin.Context) {
	rows, _ := conn.Query(context.Background(), "SELECT task_id, title FROM tasks WHERE task_id = $1", c.Param("id"))
	todoItem, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[TodoItem])
	rows.Close()
	if err != nil {
		log.Println(err)
		if errors.Is(err, pgx.ErrNoRows) {
			c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		}
		return
	}
	c.JSON(http.StatusOK, todoItem)
}

func getAllTodoItems(c *gin.Context) {
	rows, _ := conn.Query(context.Background(), "SELECT task_id, title FROM tasks")
	todoItems, err := pgx.CollectRows(rows, pgx.RowToStructByName[TodoItem])
	rows.Close()
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, todoItems)
}

func deleteTodoItem(c *gin.Context) {
	_, err := conn.Exec(context.Background(), "DELETE FROM tasks WHERE task_id = $1", c.Param("id"))
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
	router.POST("/todos/:title", createTodoItem)
	router.PUT("/todos/:id/:title", updateTodoItem)
	router.GET("/todos/:id", getTodoItemById)
	router.GET("/todos", getAllTodoItems)
	router.DELETE("/todos/:id", deleteTodoItem)
	router.Run()
}

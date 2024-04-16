package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func helloWorld(context *gin.Context) {
	context.JSON(200, gin.H{"message": "Hello World"})
}

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@db:5432/todo_list")
	if err != nil {
		fmt.Println(err)
		log.Fatalln("error connecting to the database.")
	}
	defer conn.Close(context.Background())
	router := gin.Default()
	router.GET("/helloWorld", helloWorld)
	router.Run()
}

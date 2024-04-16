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
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432")
	if err != nil {
		log.Fatalln("error connecting to the database.")
	}
	defer conn.Close(context.Background())
	router := gin.Default()
	router.GET("/helloWorld", helloWorld)
	fmt.Print("Hello World")
	router.Run("localhost:8081")
}

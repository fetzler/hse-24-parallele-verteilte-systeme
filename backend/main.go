package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func helloWorld(context *gin.Context) {
	context.JSON(200, gin.H{"message": "Hello World"})
}

func main() {
	router := gin.Default()
	router.GET("/helloWorld", helloWorld)
	fmt.Print("Hello World")
	router.Run()
}

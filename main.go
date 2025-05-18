package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go-api/handlers"
)

func main() {
	startTime := time.Now()

	r := gin.Default()

	r.GET("/users", handlers.GetUsers)
	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.GetUserByID)

	fmt.Println("Server is running on port 8080")

	elapsedTime := time.Since(startTime)
	fmt.Printf("Application started in %s\n", elapsedTime)

	r.Run(":8080")
}

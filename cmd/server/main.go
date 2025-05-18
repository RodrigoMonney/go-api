package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go-api/internal/interface/handlers"
	"go-api/internal/interface/repository"
	"go-api/internal/usecases"
)

func main() {
	startTime := time.Now()

	r := gin.Default()

	userRepo := repository.NewMemoryUserRepo()
	userUseCase := usecases.NewUserUseCase(userRepo)
	handler.NewUserHandler(r, userUseCase)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Application started in %s\n", elapsedTime)

	r.Run(":8080")
}

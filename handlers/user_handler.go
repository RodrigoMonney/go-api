package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-api/models"
	"go-api/storage"
)

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, storage.GetAllUsers())
}

func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := storage.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser := storage.CreateUser(user)
	c.JSON(http.StatusCreated, createdUser)
}
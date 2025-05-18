package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-api/internal/domain"
	"go-api/internal/usecases"
)

type UserHandler struct {
	uc *usecases.UserUseCase
}

func NewUserHandler(router *gin.Engine, uc *usecases.UserUseCase) {
	handler := &UserHandler{uc: uc}

	router.GET("/users", handler.GetAll)
	router.GET("/users/:id", handler.GetByID)
	router.POST("/users", handler.Create)
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users, _ := h.uc.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.uc.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Create(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser, _ := h.uc.CreateUser(user)
	c.JSON(http.StatusCreated, newUser)
}

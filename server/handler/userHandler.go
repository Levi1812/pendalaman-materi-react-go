package handler

import (
	"backendGolang/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

func (h *userHandler) GetAllUser(c *gin.Context) {
	users, err := h.service.GetAllUser()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *userHandler) GetUserByID(c *gin.Context) {
	paramID := c.Params.ByName("user_id")

	user, err := h.service.GetUserByID(paramID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.UserRegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	user, err := h.service.RegisterUser(input)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, user)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input user.UserLoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	user, err := h.service.LoginUser(input)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

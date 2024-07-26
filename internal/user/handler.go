package user

import (
	"net/http"

	"github.com/Just-A-NoobieDev/auction-go-server/pkg/auth"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var createUserRequest CreateUserRequest
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.UserService.RegisterUser(createUserRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

func (h *UserHandler) AuthenticateUser(c *gin.Context) {
	var authenticateUserRequest AuthenticateUserRequest
	if err := c.ShouldBindJSON(&authenticateUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserService.AuthenticateUser(authenticateUserRequest.Email, authenticateUserRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.GetString("userID")

	user, err := h.UserService.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
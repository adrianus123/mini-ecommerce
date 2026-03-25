package handler

import (
	"auth-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(u *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: u}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Email    string
		Password string
		Name     string
	}

	c.BindJSON(&req)

	err := h.usecase.Register(req.Email, req.Password, req.Name)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Register successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string
		Password string
	}

	c.BindJSON(&req)

	token, err := h.usecase.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"access_token": token})
}

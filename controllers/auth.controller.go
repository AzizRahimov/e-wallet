package controllers

import (
	"github.com/AzizRahimov/e-wallet/models"
	"github.com/AzizRahimov/e-wallet/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return AuthController{authService: authService}
}

func (h *AuthController) CreateUser(c *gin.Context) {
	var user *models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {

	}

	err = h.authService.SingUp(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "OK"})
}

func (h *AuthController) SingIn(c *gin.Context) {
	var input *models.User

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}

	token, err := h.authService.GenerateToken(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
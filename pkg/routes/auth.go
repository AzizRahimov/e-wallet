package routes

import (
	"github.com/AzizRahimov/e-wallet/pkg/controllers"
	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authRouteController controllers.AuthController
}

func NewAuthControllerRoute(authController controllers.AuthController) *AuthRouteController {
	return &AuthRouteController{authController}

}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")
	router.POST("/sing-up", rc.authRouteController.CreateUser)
	router.POST("/sing-in", rc.authRouteController.SingIn)

}

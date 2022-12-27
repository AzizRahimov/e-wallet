package routes

import "github.com/gin-gonic/gin"

type AuthRouteController struct {
	AuthRouteController controllers.AuthRouteController
}

func NewAuthControllerRoute(authController controllers.AuthRouteController) AuthRouteController {
	return AuthRouteController{authController}

}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup, ) {
	router := rg.Group("/auth")
	//router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.SignInUser)

}
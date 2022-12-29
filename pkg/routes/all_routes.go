package routes

import (
	"github.com/AzizRahimov/e-wallet/pkg/controllers"
	"github.com/AzizRahimov/e-wallet/pkg/repository"
	"github.com/AzizRahimov/e-wallet/pkg/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	server *gin.Engine
)

func RunServerAndRoutes(db *gorm.DB) {
	server = gin.Default()
	rc := server.Group("/api")
	authRep := repository.NewAuthPostgres(db)
	authService := services.NewAuthService(authRep)
	authController := controllers.NewAuthController(authService)
	authControllerRoute := NewAuthControllerRoute(authController)
	authControllerRoute.AuthRoute(rc)

	wallet := rc.Group("/wallet", authController.UserIdentity)
	walletRep := repository.NewWalletPostgres(db)
	walletService := services.NewWalletService(walletRep)
	walletController := controllers.NewWalletController(walletService)
	walletControllerRoute := NewWalletControllerRoute(walletController)
	walletControllerRoute.WalletRoute(wallet)

	server.Run(":8000")

}

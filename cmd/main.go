package main

import (
	"fmt"
	"github.com/AzizRahimov/e-wallet/controllers"
	"github.com/AzizRahimov/e-wallet/repository"
	"github.com/AzizRahimov/e-wallet/routes"
	"github.com/AzizRahimov/e-wallet/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var (
	server                *gin.Engine
	SingUpRouteController routes.AuthRouteController
)

func main() {
	execute()
	//serveApplication()

}

func execute() {
	connStr := "postgres://app:pass@localhost:5432/db?sslmode=disable"

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err, "DATABSE")
	}

	//db, err := sql.Open("postgres", connStr)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//defer func() {
	//	if cerr := db.Close(); cerr != nil {
	//		if err == nil {
	//			log.Fatal(err)
	//			return
	//		}
	//
	//	}
	//}()
	authRep := repository.NewAuthPostgres(db)
	authService := services.NewAuthService(authRep)
	authController := controllers.NewAuthController(authService)
	authControllerRoute := routes.NewAuthControllerRoute(authController)

	walletRep := repository.NewWalletPostgres(db)
	walletService := services.NewWalletService(walletRep)
	walletController := controllers.NewWalletController(walletService)
	walletControllerRoute := routes.NewWalletControllerRoute(walletController)

	server = gin.Default()

	router := server.Group("/api")
	wallet := router.Group("/wallet", authController.UserIdentity)
	router.GET("healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})
	authControllerRoute.AuthRoute(router)
	walletControllerRoute.WalletRoute(wallet)

	server.Run(":8000")
	//router.Run(":8000")
	fmt.Println("Server running on port 8000")

}

//func serveApplication() {
//
//	postService = services.NewPostService(postCollection, ctx)
//	PostController = controllers.NewPostController(postService)
//	PostRouteController = routes.NewPostControllerRoute(PostController)
//
//	//router := gin.Default()
//
//}

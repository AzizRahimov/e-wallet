package routes

import "github.com/gin-gonic/gin"

type WalletRouteController struct {
	walletRouteController controllers.WalletRouteController
}

func NewWalletControllerRoute(walletController controllers.WalletRouteController) WalletRouteController {
	return WalletRouteController{walletController}

}

func (r *WalletRouteController) WalletRoute(rg *gin.RouterGroup) {
	router := rg.Group("/wallet")

	router.POST("/check_account", controller.CheckAccount)
	router.POST("/top_up", controller.TopUp)
	router.POST("/trn_history", controller.TrnHistory)
	router.POST("/get_balance", controller.GetBalance)
}

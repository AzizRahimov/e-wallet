package routes

import (
	"github.com/AzizRahimov/e-wallet/controllers"
	"github.com/gin-gonic/gin"
)

type WalletRouteController struct {
	walletRouteController controllers.WalletController
}

func NewWalletControllerRoute(walletController controllers.WalletController) *WalletRouteController {
	return &WalletRouteController{walletController}

}

func (r *WalletRouteController) WalletRoute(rg *gin.RouterGroup) {

	//router.POST("/check_account", controller.CheckAccount)
	//router.POST("/trn_history", controller.TrnHistory)
	rg.POST("/get_balance", r.walletRouteController.GetBalance)
	rg.POST("/top_up", r.walletRouteController.TopUp)
}

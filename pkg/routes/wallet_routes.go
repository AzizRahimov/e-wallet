package routes

import (
	"github.com/AzizRahimov/e-wallet/pkg/controllers"
	"github.com/gin-gonic/gin"
)

type WalletRouteController struct {
	walletRouteController controllers.WalletController
}

func NewWalletControllerRoute(walletController controllers.WalletController) *WalletRouteController {
	return &WalletRouteController{walletController}

}

func (r *WalletRouteController) WalletRoute(rg *gin.RouterGroup) {

	rg.POST("/get_balance", r.walletRouteController.GetBalance)
	rg.POST("/top_up", r.walletRouteController.TopUp)
	rg.POST("/total_history", r.walletRouteController.TotalHistoryTrn)
	rg.POST("/check_account", r.walletRouteController.CheckAccount)
}

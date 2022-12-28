package controllers

import (
	"fmt"
	"github.com/AzizRahimov/e-wallet/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WalletController struct {
	walletService services.WalletService
}

func NewWalletController(walletService services.WalletService) WalletController {
	return WalletController{walletService: walletService}
}

func (h *WalletController) GetBalance(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}
	fmt.Println("ID", userID)
	balance, err := h.walletService.GetBalance(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balance": balance})

}

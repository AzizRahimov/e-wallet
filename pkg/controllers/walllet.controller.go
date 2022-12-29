package controllers

import (
	"encoding/json"
	"github.com/AzizRahimov/e-wallet/models"
	"github.com/AzizRahimov/e-wallet/pkg/services"
	"github.com/AzizRahimov/e-wallet/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	userIdStr := strconv.Itoa(userID)

	hash := c.GetHeader("X-Digest")

	checkHash := utils.GetSha1(userIdStr, []byte(utils.AppSettings.SecretKey.Key))

	if checkHash != hash {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "Invalid hash"})
		return
	}

	balance, err := h.walletService.GetBalance(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balance": balance})

}

func (h *WalletController) CheckAccount(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	userIdStr := strconv.Itoa(userID)

	hash := c.GetHeader("X-Digest")

	checkHash := utils.GetSha1(userIdStr, []byte(utils.AppSettings.SecretKey.Key))

	if checkHash != hash {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "Invalid hash"})
		return
	}

	account, err := h.walletService.CheckAccount(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"account": account})

}

func (h *WalletController) TopUp(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	hash := c.GetHeader("X-Digest")

	var topUp models.TopUp
	err = c.ShouldBindJSON(&topUp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	marshal, err := json.Marshal(&topUp)

	checkHash := utils.GetSha1(string(marshal), []byte(utils.AppSettings.SecretKey.Key))

	if checkHash != hash {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "Invalid hash"})
		return
	}
	topUp.ClientID = userId

	trn, err := h.walletService.TopUp(topUp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": trn})

}

func (h *WalletController) TotalHistoryTrn(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	userIdStr := strconv.Itoa(userId)

	hash := c.GetHeader("X-Digest")

	checkHash := utils.GetSha1(userIdStr, []byte(utils.AppSettings.SecretKey.Key))

	if checkHash != hash {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "Invalid hash"})
		return
	}
	trn, err := h.walletService.TotalTrn(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, trn)

}

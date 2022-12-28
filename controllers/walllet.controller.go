package controllers

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/AzizRahimov/e-wallet/models"
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
	fmt.Println(string(marshal))
	checkHash := GetSha1(string(marshal), []byte("secret"))
	fmt.Println(checkHash)

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

func GetSha1(text string, secret []byte) string {

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha1.New, secret)

	// Write Data to it
	h.Write([]byte(text))

	// Get result and encode as hexadecimal string
	hash := hex.EncodeToString(h.Sum(nil))

	return hash
}

package services

import "github.com/AzizRahimov/e-wallet/models"

type WalletService interface {
	//CheckAccount()
	TopUp(topUp models.TopUp) (trn models.Transaction, err error)
	//TrnHistory()
	GetBalance(userID int) (wallet models.Wallet, err error)
}

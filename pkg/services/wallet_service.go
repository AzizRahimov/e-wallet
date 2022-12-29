package services

import "github.com/AzizRahimov/e-wallet/models"

type WalletService interface {
	TopUp(topUp models.TopUp) (trn models.Transaction, err error)
	GetBalance(userID int) (wallet models.Wallet, err error)
	TotalTrn(userID int) (models.Transaction, error)
}

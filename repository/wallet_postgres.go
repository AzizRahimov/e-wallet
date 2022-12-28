package repository

import "github.com/AzizRahimov/e-wallet/models"

type WalletRepository interface {
	//CheckAccount()
	TopUp(wallet models.Wallet) (models.Transaction, error)
	//TrnHistory()
	GetBalance(userID int) (wallet models.Wallet, err error)
}

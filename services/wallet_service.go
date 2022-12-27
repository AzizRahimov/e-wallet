package services

import "github.com/AzizRahimov/e-wallet/models"

type WalletService interface {
	CheckAccount()
	TopUp(wallet models.Wallet)
	TrnHistory()
	GetBalance()
}

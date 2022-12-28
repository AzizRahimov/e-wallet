package repository

import (
	"github.com/AzizRahimov/e-wallet/models"
	"gorm.io/gorm"
)

type WalletRepository interface {
	//CheckAccount()
	TopUp(topUp models.TopUp) (trn models.Transaction, err error)
	//TrnHistory()
	GetBalance(userID int) (wallet models.Wallet, err error)
	GetWalletByPhoneNumber(phone string) (wallet models.Wallet, err error)
	AddTransaction(db *gorm.DB, transaction models.Transaction) (models.Transaction, error)
	GetUserByID(userID int) (user models.User, err error)
}

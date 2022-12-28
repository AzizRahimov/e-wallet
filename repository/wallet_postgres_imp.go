package repository

import (
	"errors"
	"fmt"
	"github.com/AzizRahimov/e-wallet/models"
	"gorm.io/gorm"
)

type WalletPostgresImp struct {
	db *gorm.DB
}

func NewWalletPostgres(db *gorm.DB) *WalletPostgresImp {
	return &WalletPostgresImp{db: db}
}

func (r *WalletPostgresImp) GetBalance(userID int) (wallet models.Wallet, err error) {
	query := fmt.Sprintf("SELECT id, user_id, balance FROM %q WHERE user_id = $1", "wallet")
	err = r.db.Raw(query, userID).Scan(&wallet).Error
	if err != nil {
		return wallet, err
	}
	fmt.Println("wallet.UserID", wallet.UserID)
	if wallet.UserID == 0 {
		return wallet, errors.New("account not found")
	}

	return wallet, err

}

func (r *WalletPostgresImp) TopUp(wallet models.Wallet, topUp models.TopUp) (trn models.Transaction, err error) {
	// check balance
	balance, err := r.GetBalance(wallet.UserID)
	if err != nil {
		return trn, err
	}
	if balance.Balance < wallet.Balance {
		return trn, errors.New("insufficient funds")
	}

	// check account

	tx := r.db.Begin()
	balance.Balance -= wallet.Balance
	tx.Table("wallet").Save(&balance)

	tx.Exec("INSERT into")

	return trn, nil

}

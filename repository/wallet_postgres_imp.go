package repository

import (
	"errors"
	"fmt"
	"github.com/AzizRahimov/e-wallet/models"
	"gorm.io/gorm"
	"time"
)

type WalletPostgresImp struct {
	db *gorm.DB
}

func NewWalletPostgres(db *gorm.DB) *WalletPostgresImp {
	return &WalletPostgresImp{db: db}
}

func (r *WalletPostgresImp) GetBalance(userID int) (wallet models.Wallet, err error) {
	query := fmt.Sprintf("SELECT id, user_id, balance, is_identified FROM %q WHERE user_id = $1", "wallet")
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

func (r *WalletPostgresImp) TopUp(topUp models.TopUp) (trn models.Transaction, err error) {
	// check senderWallet
	senderWallet, err := r.GetBalance(topUp.ClientID)
	if err != nil {
		return trn, err
	}
	senderUser, err := r.GetUserByID(topUp.ClientID)
	if err != nil {
		return trn, err
	}

	if senderWallet.Balance < topUp.Balance {
		return trn, errors.New("insufficient funds")
	}

	// check account

	receiverWallet, err := r.GetWalletByPhoneNumber(topUp.ReceiverPhone)
	if err != nil {
		return trn, err
	}

	if receiverWallet.IsIdentified && receiverWallet.Balance+topUp.Balance > 100000 {
		return trn, errors.New("баланс не может превышать сумму на 100 тыс")

	} else if !receiverWallet.IsIdentified && receiverWallet.Balance+topUp.Balance > 10000 {
		return trn, errors.New("баланс не может превышать сумму на 10 тыс")

	}
	tx := r.db.Begin()
	senderWallet.Balance -= topUp.Balance
	err = tx.Table("wallet").Omit("account").Save(&senderWallet).Error
	if err != nil {
		transaction, err := r.AddTransaction(r.db, models.Transaction{
			From:      senderUser.Phone,
			To:        topUp.ReceiverPhone,
			Status:    "неуспешно",
			Amount:    topUp.Balance,
			CreatedAt: time.Time{},
		})
		tx.Rollback()
		return transaction, err
	}

	receiverWallet.Balance += topUp.Balance
	err = tx.Table("wallet").Omit("account").Save(&receiverWallet).Error
	if err != nil {
		tx.Rollback()
		return trn, err
	}

	transaction, err := r.AddTransaction(tx, models.Transaction{
		From:      senderUser.Phone,
		To:        topUp.ReceiverPhone,
		Status:    "успешный",
		Amount:    topUp.Balance,
		CreatedAt: time.Time{},
	})
	tx.Commit()

	return transaction, nil

}

func (r *WalletPostgresImp) GetWalletByPhoneNumber(phone string) (wallet models.Wallet, err error) {
	query := fmt.Sprintf("SELECT w.id, w.user_id, w.balance, w.is_identified from %q w    join %q u on u.id = w.user_id WHERE  u.phone = $1", "wallet", "users")

	err = r.db.Raw(query, phone).Scan(&wallet).Error
	if err != nil {
		return wallet, err
	}

	return wallet, nil

}

func (r *WalletPostgresImp) AddTransaction(db *gorm.DB, transaction models.Transaction) (models.Transaction, error) {
	err := db.Create(&transaction).Error
	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, err

}

func (r *WalletPostgresImp) GetUserByID(userID int) (user models.User, err error) {
	query := fmt.Sprintf("SELECT  id, fio, age, phone FROM %q WHERE id = $1", "users")

	err = r.db.Raw(query, userID).Scan(&user).Error
	if user.ID == 0 {
		return models.User{}, errors.New("user not found")
	}
	if err != nil {
		return models.User{}, err
	}

	return user, nil

}

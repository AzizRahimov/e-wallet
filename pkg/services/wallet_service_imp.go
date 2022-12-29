package services

import (
	"github.com/AzizRahimov/e-wallet/models"
	"github.com/AzizRahimov/e-wallet/pkg/repository"
	"github.com/AzizRahimov/e-wallet/utils"
	"time"
)

type WalletServiceImp struct {
	repo repository.WalletRepository
}

func NewWalletService(repo repository.WalletRepository) *WalletServiceImp {
	return &WalletServiceImp{repo: repo}

}

func (s *WalletServiceImp) GetBalance(userID int) (wallet models.Wallet, err error) {
	return s.repo.GetBalance(userID)

}

func (s *WalletServiceImp) CheckAccount(userID int) (wallet models.Wallet, err error) {
	return s.repo.CheckAccount(userID)

}

func (s *WalletServiceImp) TopUp(topUp models.TopUp) (trn models.Transaction, err error) {
	return s.repo.TopUp(topUp)
}

func (s *WalletServiceImp) GetPhone(userID int) (user models.User, err error) {
	return s.repo.GetPhone(userID)

}

func (s *WalletServiceImp) TotalTrn(userID int) (models.Transaction, error) {
	currentData := time.Now().AddDate(0, 0, -time.Now().Day()+1)
	currentDataStr := currentData.Format("2006-01-02")
	userPhone, err := s.GetPhone(userID)
	if err != nil {
		return models.Transaction{}, err

	}
	totalHistoryTrnCurrentMonth, err := s.repo.GetTotalTopUpPerMonth(userPhone.Phone, currentDataStr)

	total := utils.CalculateAmount(totalHistoryTrnCurrentMonth)
	trn := models.Transaction{
		TotalAmount: total,
		Operation:   len(totalHistoryTrnCurrentMonth),
		Month:       time.Now().Month().String(),
	}

	return trn, nil

}

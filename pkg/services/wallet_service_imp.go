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

//CheckAccount - проверка на сущ аккаунта
func (s *WalletServiceImp) CheckAccount(userID int) (wallet models.Wallet, err error) {
	return s.repo.CheckAccount(userID)

}

// TopUp - Пополнение электронного кошелька
func (s *WalletServiceImp) TopUp(topUp models.TopUp) (trn models.Transaction, err error) {
	return s.repo.TopUp(topUp)
}

// GetPhone - берет данные юзера по ID
func (s *WalletServiceImp) GetPhone(userID int) (user models.User, err error) {
	return s.repo.GetPhone(userID)

}

// TotalTrn - высчитывает общую сумму и кол-во операции определенного кошелька
func (s *WalletServiceImp) TotalTrn(userID int) (models.Transaction, error) {
	//currentData := time.Now().AddDate(0, 0, -time.Now().Day()+1)
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstDayOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)

	firstDayOfMonthStr := firstDayOfMonth.Format("2006-01-02")
	lastDayOfMonthStr := lastDayOfMonth.Format("2006-01-02")

	//currentDataStr := currentData.Format("2006-01-02")

	userPhone, err := s.GetPhone(userID)
	if err != nil {
		return models.Transaction{}, err

	}
	totalHistoryTrnCurrentMonth, err := s.repo.GetTotalTopUpPerMonth(userPhone.Phone, firstDayOfMonthStr, lastDayOfMonthStr)

	total := utils.CalculateAmount(totalHistoryTrnCurrentMonth)
	trn := models.Transaction{
		TotalAmount: total,
		Operation:   len(totalHistoryTrnCurrentMonth),
		Month:       time.Now().Month().String(),
	}

	return trn, nil

}

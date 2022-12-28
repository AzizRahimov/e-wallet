package services

import (
	"github.com/AzizRahimov/e-wallet/models"
	"github.com/AzizRahimov/e-wallet/repository"
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

func (s *WalletServiceImp) TopUp(topUp models.TopUp) (trn models.Transaction, err error) {
	return s.repo.TopUp(topUp)
}

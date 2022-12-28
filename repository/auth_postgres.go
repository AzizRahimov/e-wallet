package repository

import "github.com/AzizRahimov/e-wallet/models"

type Authorization interface {
	GetUser(phone string, pin string) (*models.User, error)
	SingUp(user *models.User) error
}

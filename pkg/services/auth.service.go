package services

import "github.com/AzizRahimov/e-wallet/models"

type AuthService interface {
	ParseToken(token string) (int, error)
	GenerateToken(user models.User) (string, error)
	SingUp(user models.User) error
}

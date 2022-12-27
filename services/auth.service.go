package services

import "github.com/AzizRahimov/e-wallet/models"

type AuthService interface {
	//SignUpUser(*models.SignUpInput) (*models.DBResponse, error)
	SignInUser(user *models.User) (*models.User, error)
}

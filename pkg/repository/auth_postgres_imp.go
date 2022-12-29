package repository

import (
	"errors"
	"fmt"
	"github.com/AzizRahimov/e-wallet/models"

	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// GetUser - берет номер телефона и пин для авторизации
func (p *AuthPostgres) GetUser(phone string, pin string) (user models.User, err error) {
	query := fmt.Sprintf("SELECT id, phone FROM %q WHERE phone=$1 AND pin=$2", "users")

	err = p.db.Raw(query, phone, pin).Scan(&user).Error
	if user.ID == 0 {
		return models.User{}, errors.New("user not found")
	}

	if err != nil {
		return user, err
	}
	return user, nil
}

// SingUp - процесс регистрации юзера
func (p *AuthPostgres) SingUp(user models.User) error {
	err := p.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil

}

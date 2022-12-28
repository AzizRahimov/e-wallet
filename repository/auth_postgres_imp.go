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

func (p *AuthPostgres) GetUser(phone string, pin string) (user *models.User, err error) {
	query := fmt.Sprintf("SELECT id, phone FROM %q WHERE phone=$1 AND pin=$2", "users")

	err = p.db.Raw(query, phone, pin).Scan(&user).Error

	errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (p *AuthPostgres) SingUp(user *models.User) error {
	err := p.db.Create(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return nil

}

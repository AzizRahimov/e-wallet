package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/AzizRahimov/e-wallet/models"
	"github.com/AzizRahimov/e-wallet/pkg/repository"
	"github.com/AzizRahimov/e-wallet/utils"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

//
type AuthServiceImpl struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthServiceImpl {
	return &AuthServiceImpl{repo: repo}
}

// GenerateToken - Генируерут токен
func (s *AuthServiceImpl) GenerateToken(user models.User) (string, error) {
	user, err := s.repo.GetUser(user.Phone, generatePasswordHash(user.Pin))
	if err != nil {
		return "", err
	}
	// генерация  токена, если такой юзер существует, то генируем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(), // время жизни токена
		IssuedAt:  time.Now().Unix(),               // когда токен был создан
	},
		user.ID,
	})

	// вернем подписанного токена
	return token.SignedString([]byte(utils.AppSettings.Auth.SigningKey))

}

func (s *AuthServiceImpl) SingUp(user models.User) error {
	user.Pin = generatePasswordHash(user.Pin)
	return s.repo.SingUp(user)

}

// generatePasswordHash - хешируем пароль
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(utils.AppSettings.Auth.Salt)))

}

// ParseToken - парсить токен и возвращает его ID
func (s *AuthServiceImpl) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, errors.New("invalid signing method")
		}
		return []byte(utils.AppSettings.Auth.SigningKey), nil
	})

	if err != nil {
		return 0, nil
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClains")
	}

	return claims.UserId, nil
}

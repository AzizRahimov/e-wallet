package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/AzizRahimov/e-wallet/models"
	"github.com/AzizRahimov/e-wallet/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "dkaskdsa21312das3das"
	signingKey = "quiche#SKDJASDKA3dsa213#sH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

// db
type AuthServiceImpl struct {
	//collection *mongo.Collection
	//ctx        context.Context
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthServiceImpl {
	return &AuthServiceImpl{repo: repo}
}

//func NewAuthService(collection *mongo.Collection, ctx context.Context) AuthService {
//	return &AuthServiceImpl{collection, ctx}
//}

// GenerateToken - Генируерут токен
// а вот это ты уже добавишь в Service
func (s *AuthServiceImpl) GenerateToken(user *models.User) (string, error) {
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

	// вернем подписанный токен
	return token.SignedString([]byte(signingKey))

}

func (s *AuthServiceImpl) SingUp(user *models.User) error {
	user.Pin = generatePasswordHash(user.Pin)
	return s.repo.SingUp(user)

}

// generatePasswordHash - хешируем пароль
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))

}

// ParseToken - парсить токен и возвращает его ID
func (s *AuthServiceImpl) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
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


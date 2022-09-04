package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"library-app/entities"
	"library-app/pkg/repository"
	"time"
)

const (
	salt           = "k!2lgn[;a57^rib%jqi74a^(o;n5;7h&^bti*o(bj3;96moa*$n"
	tokenTimeLimit = 2 * time.Hour
	signingKey     = "b!$u34g%^5il9i%rn6&a74o*"
)

type AuthService struct {
	repository *repository.Repository
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repository *repository.Repository) *AuthService {
	return &AuthService{repository: repository}
}

func (as *AuthService) CreateUser(user entities.UserSignUp) (int, error) {
	user.Password = Hash(user.Password)
	userId, err := as.repository.Auth.CreateUser(user)
	if err != nil {
		return -1, err
	}
	return userId, nil
}

func (as *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := as.repository.Auth.GetUser(username, Hash(password))
	if err != nil {
		return "", err
	}
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(tokenTimeLimit)),
			IssuedAt:  jwt.NewNumericDate(now)},
		user.User_Id})

	return token.SignedString([]byte(signingKey))
}

func (as *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return -1, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return -1, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok || token.Valid {
		return -1, errors.New("access token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}

func Hash(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

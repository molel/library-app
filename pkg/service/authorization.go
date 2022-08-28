package service

import (
	"crypto/sha1"
	"fmt"
	"library-app/entities"
	"library-app/pkg/repository"
)

const (
	salt = "k!2lgn[;a57^rib%jqi74a^(o;n5;7h&^bti*o(bj3;96moa*$n"
)

type AuthService struct {
	repository *repository.Repository
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

func Hash(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

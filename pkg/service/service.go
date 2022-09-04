package service

import (
	"library-app/entities"
	"library-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user entities.UserSignUp) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Service struct {
	Auth Authorization
}

func NewService(repository *repository.Repository) *Service {
	return &Service{Auth: NewAuthService(repository)}
}

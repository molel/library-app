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

type Authors interface {
	CreateAuthor(author entities.Author) (int, error)
}

type Service struct {
	Authorization
	Authors
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository),
		Authors:       NewAuthorService(repository)}
}

package service

import (
	"library-app/entities"
	"library-app/pkg/repository"
)

type AuthorService struct {
	repository *repository.Repository
}

func NewAuthorService(repository *repository.Repository) *AuthorService {
	return &AuthorService{repository: repository}
}

func (as *AuthorService) CreateAuthor(author entities.Author) (int, error) {
	authorId, err := as.repository.CreateAuthor(author)
	if err != nil {
		return -1, err
	}
	return authorId, nil
}

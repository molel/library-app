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

func (as *AuthorService) CreateAuthor(author entities.AuthorCreate) (int, error) {
	return as.repository.CreateAuthor(author)
}

func (as *AuthorService) GetAuthors() ([]entities.AuthorUpdate, error) {
	return as.repository.GetAuthors()
}

func (as *AuthorService) GetAuthorById(id int) (entities.AuthorUpdate, error) {
	return as.repository.Authors.GetAuthorById(id)
}

func (as *AuthorService) UpdateAuthorById(id int, author entities.AuthorUpdate) error {
	return as.repository.Authors.UpdateAuthorById(id, author)
}

func (as *AuthorService) DeleteAuthorById(id int) error {
	return as.repository.Authors.DeleteAuthorById(id)
}

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

func (as *AuthorService) GetAuthors() ([]entities.Author, error) {
	authors, err := as.repository.GetAuthors()
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (as *AuthorService) GetAuthorById(id int) (entities.Author, error) {
	author, err := as.repository.Authors.GetAuthorById(id)
	if err != nil {
		return entities.Author{}, err
	}
	return author, nil
}

func (as *AuthorService) UpdateAuthorById(id int, author entities.Author) error {
	return as.repository.Authors.UpdateAuthorById(id, author)
}

func (as *AuthorService) DeleteAuthorById(id int) error {
	return as.repository.Authors.DeleteAuthorById(id)
}

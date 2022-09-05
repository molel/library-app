package service

import (
	"library-app/entities"
	"library-app/pkg/repository"
)

type BookService struct {
	repository *repository.Repository
}

func NewBookService(repository *repository.Repository) *BookService {
	return &BookService{repository: repository}
}

func (bs *BookService) CreateBook(book entities.BookCreate) (int, error) {
	bookId, err := bs.repository.Books.CreateBook(book)
	if err != nil {
		return -1, err
	}
	return bookId, nil
}

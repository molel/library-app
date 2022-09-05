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
	return bs.repository.Books.CreateBook(book)
}

func (bs *BookService) GetBooks() ([]entities.BookUpdate, error) {
	return bs.repository.Books.GetBooks()
}

func (bs *BookService) GetBookById(id int) (entities.BookUpdate, error) {
	return bs.repository.Books.GetBookById(id)
}

func (bs *BookService) UpdateBookById(id int, book entities.BookUpdate) error {
	return bs.repository.Books.UpdateBookById(id, book)
}

func (bs *BookService) DeleteBookById(id int) error {
	return bs.repository.Books.DeleteBookById(id)
}

package service

import (
	"library-app/entities"
	"library-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user entities.UserCreate) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Authors interface {
	CreateAuthor(author entities.AuthorCreate) (int, error)
	GetAuthors() ([]entities.AuthorGet, error)
	GetAuthorById(id int) (entities.AuthorGet, error)
	UpdateAuthorById(id int, author entities.AuthorUpdate) error
	DeleteAuthorById(id int) error
}

type Genres interface {
	CreateGenre(genre entities.GenreCreateAndGet) (int, error)
	GetGenres() ([]entities.GenreCreateAndGet, error)
	GetGenreById(id int) (entities.GenreCreateAndGet, error)
	UpdateGenreById(id int, genre entities.GenreUpdate) error
	DeleteGenreById(id int) error
}

type Books interface {
	CreateBook(book entities.BookCreate) (int, error)
	GetBooks() ([]entities.BookGet, error)
	GetBookById(id int) (entities.BookGet, error)
	UpdateBookById(id int, book entities.BookUpdate) error
	DeleteBookById(id int) error
}

type Service struct {
	Authorization
	Authors
	Books
	Genres
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository),
		Authors:       NewAuthorService(repository),
		Genres:        NewGenreService(repository),
		Books:         NewBookService(repository)}
}

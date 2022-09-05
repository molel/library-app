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
	CreateAuthor(author entities.AuthorCreate) (int, error)
	GetAuthors() ([]entities.AuthorUpdate, error)
	GetAuthorById(id int) (entities.AuthorCreate, error)
	UpdateAuthorById(id int, author entities.AuthorUpdate) error
	DeleteAuthorById(id int) error
}

type Genres interface {
	CreateGenre(genre entities.GenreCreate) (int, error)
	GetGenres() ([]entities.GenreCreate, error)
	GetGenreById(id int) (entities.GenreCreate, error)
	UpdateGenreById(id int, genre entities.GenreUpdate) error
	DeleteGenreById(id int) error
}

type Books interface {
	CreateBook(book entities.BookCreate) (int, error)
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

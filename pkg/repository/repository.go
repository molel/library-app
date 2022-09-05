package repository

import (
	"github.com/jmoiron/sqlx"
	"library-app/entities"
)

const (
	usersTableName   = "users"
	authorsTableName = "authors"
	genresTableName  = "genres"
	booksTableName   = "books"
)

type Authorization interface {
	CreateUser(up entities.UserSignUp) (int, error)
	GetUser(username, password string) (entities.User, error)
	GetUserId(username, password string) (int, error)
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

type Repository struct {
	Authorization
	Authors
	Genres
	Books
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
		Authors:       NewAuthorDB(db),
		Genres:        NewGenreDB(db),
		Books:         NewBookDB(db)}
}

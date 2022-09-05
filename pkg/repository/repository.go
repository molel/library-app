package repository

import (
	"github.com/jmoiron/sqlx"
	"library-app/entities"
)

const (
	usersTableName   = "users"
	authorsTableName = "authors"
)

type Authorization interface {
	CreateUser(up entities.UserSignUp) (int, error)
	GetUser(username, password string) (entities.User, error)
	GetUserId(username, password string) (int, error)
}

type Authors interface {
	CreateAuthor(author entities.Author) (int, error)
	GetAuthors() ([]entities.Author, error)
	GetAuthorById(id int) (entities.Author, error)
	UpdateAuthorById(id int, author entities.Author) error
	DeleteAuthorById(id int) error
}

type Repository struct {
	Authorization
	Authors
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
		Authors:       NewAuthorDB(db)}
}

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
}

type Authors interface {
	CreateAuthor(author entities.Author) (int, error)
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

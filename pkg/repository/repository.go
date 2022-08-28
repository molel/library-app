package repository

import (
	"github.com/jmoiron/sqlx"
	"library-app/entities"
)

const (
	usersTableName = "users"
)

type Authorization interface {
	CreateUser(up entities.UserSignUp) (int, error)
}

type Repository struct {
	Auth Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Auth: NewAuthDB(db)}
}

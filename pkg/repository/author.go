package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"library-app/entities"
)

type AuthorDB struct {
	*sqlx.DB
}

func NewAuthorDB(db *sqlx.DB) *AuthorDB {
	return &AuthorDB{db}
}

func (db *AuthorDB) CreateAuthor(author entities.Author) (int, error) {
	var authorId int
	query := fmt.Sprintf("INSERT INTO %s(name, surname, description) VALUES($1, $2, $3) RETURNING author_id;", authorsTableName)
	if err := db.QueryRow(query, author.Name, author.Surname, author.Description).Scan(&authorId); err != nil {
		return -1, err
	}
	return authorId, nil
}

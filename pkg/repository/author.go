package repository

import (
	"errors"
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

func (db *AuthorDB) CreateAuthor(author entities.AuthorCreate) (int, error) {
	var authorId int
	query := fmt.Sprintf("INSERT INTO %s(name, surname, description) VALUES($1, $2, $3) RETURNING author_id;", authorsTableName)
	if err := db.QueryRow(query, author.Name, author.Surname, author.Description).Scan(&authorId); err != nil {
		return -1, err
	}
	return authorId, nil
}

func (db *AuthorDB) GetAuthors() ([]entities.AuthorGet, error) {
	var author entities.AuthorGet
	authors := make([]entities.AuthorGet, 0)
	query := fmt.Sprintf("SELECT author_id AS authorId, name, surname, description FROM %s", authorsTableName)
	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&author)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return authors, nil
}

func (db *AuthorDB) GetAuthorById(id int) (entities.AuthorGet, error) {
	if exist := Exists(db.DB, authorsTableName, "author_id", id); !exist {
		return entities.AuthorGet{}, errors.New("there is no authors with such id")
	}
	var author entities.AuthorGet
	query := fmt.Sprintf("SELECT author_id AS authorId, name, surname, description FROM %s WHERE author_id = $1", authorsTableName)
	if err := db.Get(&author, query, id); err != nil {
		return entities.AuthorGet{}, err
	}
	return author, nil

}

func (db *AuthorDB) UpdateAuthorById(id int, author entities.AuthorUpdate) error {
	if exist := Exists(db.DB, authorsTableName, "author_id", id); !exist {
		return errors.New("there is no authors with such id")
	}
	fields, values, err := getUpdateArgs(author)
	if err != nil {
		return err
	}
	values = append(values, id)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE author_id = $%d", authorsTableName, fields, len(values))
	_, err = db.Exec(query, values...)
	return err

}

func (db *AuthorDB) DeleteAuthorById(id int) error {
	if exist := Exists(db.DB, authorsTableName, "author_id", id); !exist {
		return errors.New("there is no authors with such id")
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE author_id = $1", authorsTableName)
	_, err := db.Exec(query, id)
	return err
}

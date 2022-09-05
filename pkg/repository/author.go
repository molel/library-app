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

func (db *AuthorDB) GetAuthors() ([]entities.Author, error) {
	var author entities.Author
	authors := make([]entities.Author, 0)
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
	return authors, nil
}

func (db *AuthorDB) GetAuthorById(id int) (entities.Author, error) {
	var author entities.Author
	query := fmt.Sprintf("SELECT author_id AS authorId, name, surname, description FROM %s WHERE author_id = $1", authorsTableName)
	if err := db.Get(&author, query, id); err != nil {
		return entities.Author{}, err
	}
	return author, nil
}

func (db *AuthorDB) UpdateAuthorById(id int, author entities.Author) error {
	query := fmt.Sprintf("UPDATE %s SET name = $2, surname = $3, description = $4 WHERE author_id = $1", authorsTableName)
	_, err := db.Exec(query, id, author.Name, author.Surname, author.Description.String)
	return err
}

func (db *AuthorDB) DeleteAuthorById(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE author_id = $1", authorsTableName)
	_, err := db.Exec(query, id)
	return err
}

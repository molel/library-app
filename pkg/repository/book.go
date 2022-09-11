package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"library-app/entities"
)

type BookDB struct {
	*sqlx.DB
}

func NewBookDB(db *sqlx.DB) *BookDB {
	return &BookDB{db}
}

func (db *BookDB) CreateBook(book entities.BookCreate) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(name, description, genre_id, author_id) VALUES($1, $2, $3, $4) RETURNING id", booksTableName)
	if err := db.QueryRow(query, book.Name, book.Description, book.GenreId, book.AuthorId).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (db *BookDB) GetBooks() ([]entities.BookGet, error) {
	var book entities.BookGet
	books := make([]entities.BookGet, 0)
	query := fmt.Sprintf("SELECT id, name, description, genre_id AS genreId, author_id AS authorId FROM %s", booksTableName)
	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&book)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return books, nil
}

func (db *BookDB) GetBookById(id int) (entities.BookGet, error) {
	if exist := Exists(db.DB, booksTableName, []string{"id"}, []interface{}{id}); !exist {
		return entities.BookGet{}, errors.New("there is no book with such id")
	}
	var book entities.BookGet
	query := fmt.Sprintf("SELECT id, name, description, genre_id AS genreId, author_id AS authorId FROM %s WHERE id = $1", booksTableName)
	if err := db.Get(&book, query, id); err != nil {
		return entities.BookGet{}, err
	}
	return book, nil
}

func (db *BookDB) UpdateBookById(id int, book entities.BookUpdate) error {
	if exist := Exists(db.DB, booksTableName, []string{"id"}, []interface{}{id}); !exist {
		return errors.New("there is no book with such id")
	}
	fields, values, err := getUpdateArgs(book)
	if err != nil {
		return err
	}
	values = append(values, id)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", booksTableName, fields, len(values))
	_, err = db.Exec(query, values...)
	return err
}

func (db *BookDB) DeleteBookById(id int) error {
	if exist := Exists(db.DB, booksTableName, []string{"id"}, []interface{}{id}); !exist {
		return errors.New("there is no book with such id")
	}
	// TODO don't forget to set list_items ON DELETE CASCADE
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", booksTableName)
	_, err := db.Exec(query, id)
	return err
}

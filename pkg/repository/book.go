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
	var bookId int
	query := fmt.Sprintf("INSERT INTO %s(name, description, genre_id, author_id) VALUES($1, $2, $3, $4) RETURNING book_id", booksTableName)
	if err := db.QueryRow(query, book.Name, book.Description, book.GenreId, book.AuthorId).Scan(&bookId); err != nil {
		return -1, err
	}
	return bookId, nil
}

func (db *BookDB) GetBooks() ([]entities.BookGet, error) {
	var book entities.BookGet
	books := make([]entities.BookGet, 0)
	query := fmt.Sprintf("SELECT book_id AS bookId, name, description, genre_id AS genreId, author_id AS authorId FROM %s", booksTableName)
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
	if exist := Exists(db.DB, booksTableName, "book_id", id); !exist {
		return entities.BookGet{}, errors.New("there is no books with such id")
	}
	var book entities.BookGet
	query := fmt.Sprintf("SELECT book_id AS bookId, name, description, genre_id AS genreId, author_id AS authorId FROM %s WHERE book_id = $1", booksTableName)
	if err := db.Get(&book, query, id); err != nil {
		return entities.BookGet{}, err
	}
	return book, nil
}

func (db *BookDB) UpdateBookById(id int, book entities.BookUpdate) error {
	if exist := Exists(db.DB, booksTableName, "book_id", id); !exist {
		return errors.New("there is no books with such id")
	}
	fields, values, err := getUpdateArgs(book)
	if err != nil {
		return err
	}
	values = append(values, id)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE book_id = $%d", booksTableName, fields, len(values))
	_, err = db.Exec(query, values...)
	return err
}

func (db *BookDB) DeleteBookById(id int) error {
	if exist := Exists(db.DB, booksTableName, "book_id", id); !exist {
		return errors.New("there is no books with such id")
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE book_id = $1", booksTableName)
	_, err := db.Exec(query, id)
	return err
}

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
	if err := db.QueryRow(query, book.Name, book.Description.String, book.GenreId, book.AuthorId).Scan(&bookId); err != nil {
		return -1, err
	}
	return bookId, nil
}

func (db *BookDB) GetBooks() ([]entities.BookUpdate, error) {
	var book entities.BookUpdate
	books := make([]entities.BookUpdate, 0)
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

func (db *BookDB) GetBookById(id int) (entities.BookUpdate, error) {
	if exist := db.Exist(id); exist {
		var book entities.BookUpdate
		query := fmt.Sprintf("SELECT book_id AS bookId, name, description, genre_id AS genreId, author_id AS authorId FROM %s WHERE book_id = $1", booksTableName)
		if err := db.Get(&book, query, id); err != nil {
			return entities.BookUpdate{}, err
		}
		return book, nil
	} else {
		return entities.BookUpdate{}, errors.New("there is no books with such id")
	}
}

func (db *BookDB) UpdateBookById(id int, book entities.BookUpdate) error {
	if exist := db.Exist(id); exist {
		query := fmt.Sprintf("UPDATE %s SET name = $2, description = $3, genre_id = $4, author_id = $5 WHERE book_id = $1", booksTableName)
		_, err := db.Exec(query, id, book.Name, book.Description.String, book.GenreId, book.AuthorId)
		return err
	} else {
		return errors.New("there is no books with such id")
	}
}

func (db *BookDB) DeleteBookById(id int) error {
	if exist := db.Exist(id); exist {
		query := fmt.Sprintf("DELETE FROM %s WHERE book_id = $1", booksTableName)
		_, err := db.Exec(query, id)
		return err
	} else {
		return errors.New("there is no books with such id")
	}
}

func (db *BookDB) Exist(id int) bool {
	var exist bool
	query := "SELECT EXISTS(SELECT 1 FROM books WHERE book_id = $1)"
	if err := db.QueryRow(query, id).Scan(&exist); err != nil {
		return false
	}
	return exist
}

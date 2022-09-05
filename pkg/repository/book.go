package repository

import (
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
	query := fmt.Sprintf("INSERT INTO %s(name, description, genre_id, author_id) ", booksTableName)
	if err := db.QueryRow(query, book.Name, book.Description, book.GenreId, book.AuthorId).Scan(&bookId); err != nil {
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
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&book)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (db *BookDB) GetBookById(id int) (entities.BookUpdate, error) {
	var book entities.BookUpdate
	query := fmt.Sprintf("SELECT book_id AS bookId, name, description, genre_id AS genreId, author_id AS authorId FROM %s WHERE book_id = $1", booksTableName)
	if err := db.Get(&book, query, id); err != nil {
		return entities.BookUpdate{}, err
	}
	return book, nil
}

func (db *BookDB) UpdateBookById(id int, book entities.BookUpdate) error {
	query := fmt.Sprintf("UPDATE %s SET name = $2, description = $3, genre_id = $4, author_id = $5 WHERE book_id = $1", booksTableName)
	_, err := db.Exec(query, id, book.Name, book.Description.String, book.GenreId, book.AuthorId)
	return err
}

func (db *BookDB) DeleteBookById(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE book_id = $1", booksTableName)
	_, err := db.Exec(query, id)
	return err
}

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

func (db BookDB) CreateBook(book entities.BookCreate) (int, error) {
	var bookId int
	query := fmt.Sprintf("INSERT INTO %s(name, description, genre_id, author_id) ", booksTableName)
	if err := db.QueryRow(query, book.Name, book.Description, book.GenreId, book.AuthorId).Scan(&bookId); err != nil {
		return -1, err
	}
	return bookId, nil
}

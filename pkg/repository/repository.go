package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"library-app/entities"
	"reflect"
	"strings"
)

const (
	usersTableName     = "users"
	authorsTableName   = "authors"
	genresTableName    = "genres"
	booksTableName     = "books"
	listsTableName     = "lists"
	listItemsTableName = "list_items"
)

type Authorization interface {
	CreateUser(up entities.UserCreate) (int, error)
	GetUser(username, password string) (entities.UserGet, error)
	GetUserId(username, password string) (int, error)
}

type Authors interface {
	CreateAuthor(author entities.AuthorCreate) (int, error)
	GetAuthors() ([]entities.AuthorGet, error)
	GetAuthorById(id int) (entities.AuthorGet, error)
	UpdateAuthorById(id int, author entities.AuthorUpdate) error
	DeleteAuthorById(id int) error
}

type Genres interface {
	CreateGenre(genre entities.GenreCreateAndGet) (int, error)
	GetGenres() ([]entities.GenreCreateAndGet, error)
	GetGenreById(id int) (entities.GenreCreateAndGet, error)
	UpdateGenreById(id int, genre entities.GenreUpdate) error
	DeleteGenreById(id int) error
}

type Books interface {
	CreateBook(book entities.BookCreate) (int, error)
	GetBooks() ([]entities.BookGet, error)
	GetBookById(id int) (entities.BookGet, error)
	UpdateBookById(id int, book entities.BookUpdate) error
	DeleteBookById(id int) error
}

type Lists interface {
	CreateList(userId int, create entities.ListCreate) (int, error)
	GetLists(userId int) ([]entities.ListGet, error)
	GetListById(userId, id int) (entities.ListGetWithItems, error)
	UpdateListById(userId, id int, list entities.ListUpdate) error
	DeleteListById(userId, id int) error
}

type ListItems interface {
	CreateListItem(userId, listId int, listItem entities.ListItemCreate) error
	UpdateListItemById(userId, listId, bookId int, listItem entities.ListItemUpdate) error
	DeleteListItemById(userId, listId, bookId int) error
}

type Repository struct {
	Authorization
	Authors
	Genres
	Books
	Lists
	ListItems
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
		Authors:       NewAuthorDB(db),
		Genres:        NewGenreDB(db),
		Books:         NewBookDB(db),
		Lists:         NewListDB(db),
		ListItems:     NewListItemDB(db)}
}

func Exists(db *sqlx.DB, table string, columns []string, values []interface{}) bool {
	var exist bool
	args := make([]string, 0)
	for i := 0; i < len(columns); i++ {
		args = append(args, fmt.Sprintf("%s = $%d", columns[i], i+1))
	}
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s)", table, strings.Join(args, " AND "))
	if err := db.QueryRow(query, values...).Scan(&exist); err != nil {
		return false
	}
	return exist
}

func getUpdateArgs(entity interface{}) (string, []interface{}, error) {
	v := reflect.ValueOf(entity)
	if v.NumField() == 0 {
		return "", nil, errors.New("empty update entity")
	}
	fields := make([]string, 0)
	values := make([]interface{}, 0)
	n := 1
	for i := 0; i < v.NumField(); i++ {
		if !v.Field(i).IsNil() {
			fields = append(fields, fmt.Sprintf("%s = $%d", v.Type().Field(i).Name, n))
			values = append(values, v.Field(i).Interface())
			n++
		}
	}
	return strings.Join(fields, ", "), values, nil
}

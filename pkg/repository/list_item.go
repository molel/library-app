package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"library-app/entities"
)

type ListItemDB struct {
	*sqlx.DB
}

func NewListItemDB(db *sqlx.DB) *ListItemDB {
	return &ListItemDB{db}
}

func (db *ListItemDB) CreateListItem(userId, listId int, listItem entities.ListItemCreate) error {
	if exist := Exists(db.DB, listsTableName, []string{"id", "user_id"}, []interface{}{listId, userId}); !exist {
		return errors.New("there is no list with such id")
	}
	query := fmt.Sprintf("INSERT INTO %s(list_id, book_id, status) VALUES($1, $2, $3)", listItemsTableName)
	_, err := db.Exec(query, listId, listItem.BookId, listItem.Status)
	return err
}

func (db *ListItemDB) UpdateListItemById(userId, listId, bookId int, listItem entities.ListItemUpdate) error {
	if exist := Exists(db.DB, listsTableName, []string{"id", "user_id"}, []interface{}{listId, userId}); !exist {
		return errors.New("there is no list with such id")
	}
	fields, values, err := getUpdateArgs(listItem)
	if err != nil {
		return err
	}
	values = append(values, listId, bookId)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE list_id = $%d AND book_id = $%d", listItemsTableName, fields, len(values)-1, len(values))
	_, err = db.Exec(query, values...)
	return err
}

func (db *ListItemDB) DeleteListItemById(userId, listId, bookId int) error {
	if exist := Exists(db.DB, listsTableName, []string{"id", "user_id"}, []interface{}{listId, userId}); !exist {
		return errors.New("there is no list with such id")
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE list_id = $1 AND book_id = $2", listItemsTableName)
	_, err := db.Exec(query, listId, bookId)
	return err
}

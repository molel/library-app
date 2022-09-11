package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"library-app/entities"
)

type ListDB struct {
	*sqlx.DB
}

func NewListDB(db *sqlx.DB) *ListDB {
	return &ListDB{db}
}

func (db *ListDB) CreateList(userId int, list entities.ListCreate) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(title, user_id) VALUES($1, $2) RETURNING id", listsTableName)
	if err := db.QueryRow(query, list.Title, userId).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (db *ListDB) GetLists(userId int) ([]entities.ListGet, error) {
	var list entities.ListGet
	lists := make([]entities.ListGet, 0)
	query := fmt.Sprintf("SELECT id, title, user_id AS userId FROM %s", listsTableName)
	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&list)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return lists, nil
}

func (db *ListDB) GetListById(userId, id int) (entities.ListGet, error) {
	if exist := Exists(db.DB, listsTableName, []string{"id"}, []interface{}{id}); !exist {
		return entities.ListGet{}, errors.New("there is no list with such id")
	}
	var list entities.ListGet
	query := fmt.Sprintf("SELECT id, title, user_id AS userId FROM %s WHERE id = $1 AND user_id = $2", listsTableName)
	if err := db.Get(&list, query, id, userId); err != nil {
		return entities.ListGet{}, err
	}
	return list, nil
}

func (db *ListDB) UpdateListById(userId, id int, list entities.ListUpdate) error {
	if exist := Exists(db.DB, listsTableName, []string{"id"}, []interface{}{id}); !exist {
		return errors.New("there is no list with such id")
	}
	fields, values, err := getUpdateArgs(list)
	if err != nil {
		return err
	}
	values = append(values, id, userId)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d AND user_id = $%d", listsTableName, fields, len(values)-1, len(values))
	_, err = db.Exec(query, values...)
	return err
}

func (db *ListDB) DeleteListById(userId, id int) error {
	if exist := Exists(db.DB, listsTableName, []string{"id", "user_id"}, []interface{}{id, userId}); !exist {
		return errors.New("there is no list with such id")
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 AND user_id = $2", listsTableName)
	_, err := db.Exec(query, id, userId)
	return err
}

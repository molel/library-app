package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"library-app/entities"
)

type AuthDB struct {
	*sqlx.DB
}

func NewAuthDB(db *sqlx.DB) *AuthDB {
	return &AuthDB{db}
}

func (db *AuthDB) CreateUser(user entities.UserCreate) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(username, password) VALUES($1, $2) RETURNING id", usersTableName)
	if err := db.QueryRow(query, user.Username, user.Password).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (db *AuthDB) GetUserId(username, password string) (int, error) {
	if exist := Exists(db.DB, usersTableName, "username", username); !exist {
		return -1, errors.New("there is no users with such username")
	}
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password = $2;", usersTableName)
	err := db.QueryRow(query, username, password).Scan(&id)
	return id, err
}

func (db *AuthDB) GetUser(username, password string) (entities.UserGet, error) {
	var user entities.UserGet
	query := fmt.Sprintf("SELECT * FROM %s WHERE username = $1 AND password = $2;", usersTableName)
	err := db.Get(&user, query, username, password)
	return user, err
}

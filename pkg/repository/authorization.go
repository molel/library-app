package repository

import (
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

func (db *AuthDB) CreateUser(user entities.UserSignUp) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s(username, password) VALUES($1, $2) RETURNING user_id", usersTableName)
	var userId int
	if err := db.QueryRow(query, user.Username, user.Password).Scan(&userId); err != nil {
		return -1, err
	}
	return userId, nil
}

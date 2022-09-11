package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"library-app/entities"
)

type GenreDB struct {
	*sqlx.DB
}

func NewGenreDB(db *sqlx.DB) *GenreDB {
	return &GenreDB{db}
}

func (db *GenreDB) CreateGenre(genre entities.GenreCreateAndGet) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(id, name) VALUES($1, $2) RETURNING id", genresTableName)
	if err := db.QueryRow(query, genre.Id, genre.Name).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (db *GenreDB) GetGenres() ([]entities.GenreCreateAndGet, error) {
	var genre entities.GenreCreateAndGet
	genres := make([]entities.GenreCreateAndGet, 0)
	query := fmt.Sprintf("SELECT id, name FROM %s", genresTableName)
	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&genre)
		if err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return genres, nil
}

func (db *GenreDB) GetGenreById(id int) (entities.GenreCreateAndGet, error) {
	if exist := Exists(db.DB, genresTableName, []string{"id"}, []interface{}{id}); !exist {
		return entities.GenreCreateAndGet{}, errors.New("there is no genre with such id")
	}
	var genre entities.GenreCreateAndGet
	query := fmt.Sprintf("SELECT id AS id, name FROM %s WHERE id = $1", genresTableName)
	if err := db.Get(&genre, query, id); err != nil {
		return entities.GenreCreateAndGet{}, err
	}
	return genre, nil
}

func (db *GenreDB) UpdateGenreById(id int, genre entities.GenreUpdate) error {
	if exist := Exists(db.DB, genresTableName, []string{"id"}, []interface{}{id}); !exist {
		return errors.New("there is no genre with such id")
	}
	fields, values, err := getUpdateArgs(genre)
	if err != nil {
		return err
	}
	values = append(values, id)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", genresTableName, fields, len(values))
	_, err = db.Exec(query, values...)
	return err

}

func (db *GenreDB) DeleteGenreById(id int) error {
	if exist := Exists(db.DB, genresTableName, []string{"id"}, []interface{}{id}); !exist {
		return errors.New("there is no genre with such id")
	}
	if exist := Exists(db.DB, booksTableName, []string{"genre_id"}, []interface{}{id}); exist {
		return errors.New("cannot delete genre: genre has 1 or more dependencies")
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", genresTableName)
	_, err := db.Exec(query, id)
	return err
}

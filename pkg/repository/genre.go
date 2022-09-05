package repository

import (
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

func (db *GenreDB) CreateGenre(genre entities.GenreCreate) (int, error) {
	var genreId int
	query := fmt.Sprintf("INSERT INTO %s(genre_id, name) VALUES($1, $2) RETURNING genre_id", genresTableName)
	if err := db.QueryRow(query, genre.GenreId, genre.Name).Scan(&genreId); err != nil {
		return -1, err
	}
	return genreId, nil
}

func (db *GenreDB) GetGenres() ([]entities.GenreCreate, error) {
	var genre entities.GenreCreate
	genres := make([]entities.GenreCreate, 0)
	query := fmt.Sprintf("SELECT genre_id AS genreId, name FROM %s", genresTableName)
	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&genre)
		if err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}
	return genres, nil
}

func (db *GenreDB) GetGenreById(id int) (entities.GenreCreate, error) {
	var genre entities.GenreCreate
	query := fmt.Sprintf("SELECT genre_id AS genreId, name FROM %s WHERE genre_id = $1", genresTableName)
	if err := db.Get(&genre, query, id); err != nil {
		return entities.GenreCreate{}, err
	}
	return genre, nil
}

func (db *GenreDB) UpdateGenreById(id int, genre entities.GenreUpdate) error {
	query := fmt.Sprintf("UPDATE %s SET name = $2 WHERE genre_id = $1", genresTableName)
	_, err := db.Exec(query, id, genre.Name)
	return err
}

func (db *GenreDB) DeleteGenreById(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE genre_id = $1", genresTableName)
	_, err := db.Exec(query, id)
	return err
}

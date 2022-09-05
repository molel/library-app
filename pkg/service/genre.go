package service

import (
	"library-app/entities"
	"library-app/pkg/repository"
)

type GenreService struct {
	repository *repository.Repository
}

func NewGenreService(repository *repository.Repository) *GenreService {
	return &GenreService{repository: repository}
}

func (gs *GenreService) CreateGenre(genre entities.GenreCreate) (int, error) {
	return gs.repository.Genres.CreateGenre(genre)
}

func (gs *GenreService) GetGenres() ([]entities.GenreCreate, error) {
	return gs.repository.Genres.GetGenres()
}

func (gs *GenreService) GetGenreById(id int) (entities.GenreCreate, error) {
	return gs.repository.Genres.GetGenreById(id)
}

func (gs *GenreService) UpdateGenreById(id int, genre entities.GenreUpdate) error {
	return gs.repository.UpdateGenreById(id, genre)
}

func (gs *GenreService) DeleteGenreById(id int) error {
	return gs.repository.Genres.DeleteGenreById(id)
}

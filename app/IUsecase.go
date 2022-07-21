package app

import (
	"github.com/DarkSoul94/film-rest/models"
	"github.com/google/uuid"
)

// Usecase ...
type Usecase interface {
	GetAllFilms(sortType string) ([]models.Film, error)
	CreateFilm(film models.Film) error
	GetFilm(id uuid.UUID) (models.Film, error)
}

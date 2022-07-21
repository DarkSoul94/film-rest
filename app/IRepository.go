package app

import (
	"github.com/DarkSoul94/film-rest/models"
	"github.com/google/uuid"
)

// Repository ...
type Repository interface {
	GetAllFilms() ([]models.Film, error)
	AddFilm(film models.Film) error
	GetFilmById(id uuid.UUID) (models.Film, error)
}

package app

import (
	"github.com/DarkSoul94/film-rest/models"
)

// Usecase ...
type Usecase interface {
	GetAllFilms() ([]models.Film, error)
	CreateFilm(film models.Film) error
}

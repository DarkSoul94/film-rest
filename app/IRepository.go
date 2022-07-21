package app

import "github.com/DarkSoul94/film-rest/models"

// Repository ...
type Repository interface {
	GetAllFilms() ([]models.Film, error)
	AddFilm(film models.Film) error
}

package app

import "github.com/DarkSoul94/film-rest/models"

// Repository ...
type Repository interface {
	AddFilm(film models.Film) error
}

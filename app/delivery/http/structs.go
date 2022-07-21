package http

import (
	"errors"
	"time"

	"github.com/DarkSoul94/film-rest/models"
	"github.com/google/uuid"
)

type newFilm struct {
	Title       string `json:"title"`
	ReleaseDate string `json:"releaseDate"`
}

func (h *Handler) toModelFilm(film newFilm) (models.Film, error) {
	release, err := time.Parse(`02.01.2006T15:04:05-0700`, film.ReleaseDate)
	if err != nil {
		return models.Film{}, err
	}

	if len(film.Title) == 0 {
		return models.Film{}, errors.New("title is empty")
	}

	return models.Film{
		Title:       film.Title,
		ReleaseDate: release,
	}, nil
}

type outFilm struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"releaseDate"`
}

func (h *Handler) toOutFilm(film models.Film) outFilm {
	return outFilm{
		ID:          film.ID,
		Title:       film.Title,
		ReleaseDate: film.ReleaseDate,
	}
}

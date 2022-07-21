package usecase

import (
	"sort"

	"github.com/DarkSoul94/film-rest/app"
	"github.com/DarkSoul94/film-rest/models"
	"github.com/google/uuid"
)

const (
	Asc  = "releaseDate"
	Desc = "-releaseDate"
)

// Usecase ...
type usecase struct {
	repo app.Repository
}

// NewUsecase ...
func NewUsecase(repo app.Repository) app.Usecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) CreateFilm(film models.Film) error {
	return u.repo.AddFilm(film)
}

func (u *usecase) GetAllFilms(sortType string) ([]models.Film, error) {
	var (
		films []models.Film = make([]models.Film, 0)
		err   error
	)

	films, err = u.repo.GetAllFilms()
	if err != nil {
		return nil, err
	}

	switch sortType {
	case Asc:
		sort.Slice(films, func(i, j int) bool {
			return films[i].ReleaseDate.Before(films[j].ReleaseDate)
		})
	case Desc:
		sort.Slice(films, func(i, j int) bool {
			return films[i].ReleaseDate.After(films[j].ReleaseDate)
		})
	}

	return films, nil
}

func (u *usecase) GetFilm(id uuid.UUID) (models.Film, error) {
	return u.repo.GetFilmById(id)
}

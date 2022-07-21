package usecase

import (
	"github.com/DarkSoul94/film-rest/app"
	"github.com/DarkSoul94/film-rest/models"
	"github.com/google/uuid"
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

func (u *usecase) GetAllFilms() ([]models.Film, error) {
	return u.repo.GetAllFilms()
}

func (u *usecase) GetFilm(id uuid.UUID) (models.Film, error) {
	return u.repo.GetFilmById(id)
}

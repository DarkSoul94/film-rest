package usecase

import (
	"time"

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
	film.ID = uuid.New()
	film.CreatedAt = time.Now()

	return u.repo.AddFilm(film)
}

func (u *usecase) GetAllFilms() ([]models.Film, error) {
	return u.repo.GetAllFilms()
}

package app

import (
	"context"

	"github.com/DarkSoul94/film-rest/models"
)

// Usecase ...
type Usecase interface {
	CreateFilm(ctx context.Context, film models.Film) error
}

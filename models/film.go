package models

import (
	"time"

	"github.com/google/uuid"
)

type Film struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	CreatedAt   time.Time `json:"created_at"`
}

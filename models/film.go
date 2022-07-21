package models

import (
	"time"

	"github.com/google/uuid"
)

type Film struct {
	ID          uuid.UUID
	Title       string
	ReleaseDate time.Time
	CreatedAt   time.Time
}

package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UUID        uuid.UUID
	Email       string
	AccessToken string
}

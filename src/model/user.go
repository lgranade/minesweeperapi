package model

import (
	"time"

	"github.com/google/uuid"
)

// User represents the public model for a user
type User struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

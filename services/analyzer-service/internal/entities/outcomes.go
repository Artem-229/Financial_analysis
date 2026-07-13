package entities

import "github.com/google/uuid"

type Outcome struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	Analysis string
}

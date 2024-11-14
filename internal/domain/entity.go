package domain

import "github.com/google/uuid"

type IEntity interface {
	GetIdentifier() uuid.UUID
}

package entities

import "github.com/google/uuid"

type ID = uuid.UUID

func NewId () ID {
	return ID(uuid.New())
}

func parseID (s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
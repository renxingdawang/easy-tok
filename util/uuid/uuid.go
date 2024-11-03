package uuid

import "github.com/google/uuid"

func GenUUID() uuid.UUID {
	id := uuid.New()
	return id
}

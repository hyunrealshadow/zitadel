package id

import (
	"github.com/google/uuid"
)

type uuidGenerator struct {
}

func (s *uuidGenerator) Next() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

var _uuidGenerator Generator

// UUIDGenerator creates a new id generator
func UUIDGenerator() Generator {
	if _uuidGenerator == nil {
		generator := Generator(&uuidGenerator{})
		_uuidGenerator = generator
	}
	return _uuidGenerator
}

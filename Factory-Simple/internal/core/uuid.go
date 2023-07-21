package core

import (
	"github.com/google/uuid"
)

// 二次封裝UUID套件
type UUID struct {
	uuid uuid.UUID
}

func NewUUID() UUID {
	return UUID{
		uuid: uuid.New(),
	}
}

func ParseUUID(id string) (UUID, error) {
	if len(id) < 1 {
		return UUID{uuid: uuid.Nil}, nil
	}

	result, err := uuid.Parse(id)
	if err != nil {
		return UUID{uuid: uuid.Nil}, err
	}

	return UUID{uuid: result}, nil
}

func (u UUID) ToString() string {
	return u.uuid.String()
}

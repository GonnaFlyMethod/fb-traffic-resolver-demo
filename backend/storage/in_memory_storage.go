package storage

import (
	"backend/common"
	"github.com/google/uuid"
)

type InMemoryUserStorage struct {
	data []common.UserWithID
}

func NewInMemoryStorage() *InMemoryUserStorage {
	return &InMemoryUserStorage{data: []common.UserWithID{}}
}

func (s *InMemoryUserStorage) GetAllUsers() []common.UserWithID {
	return s.data
}

func (s *InMemoryUserStorage) CreateNewUser(user common.UserWithoutID) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	newUser := common.UserWithID{
		Id:      id.String(),
		Email:   user.Email,
		Name:    user.Name,
		Surname: user.Surname,
	}
	s.data = append(s.data, newUser)
	return nil
}

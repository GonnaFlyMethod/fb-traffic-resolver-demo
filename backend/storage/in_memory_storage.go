package storage

import (
	"backend/common"
	"github.com/google/uuid"
)

type InMemoryUserStorage struct {
	data []*common.UserWithID
}

func NewInMemoryStorage() *InMemoryUserStorage {
	return &InMemoryUserStorage{data: []*common.UserWithID{}}
}

func (s *InMemoryUserStorage) GetAllUsers() []*common.UserWithID {
	return s.data
}

func (s *InMemoryUserStorage) CreateNewUser(user common.UserWithoutID) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	newUser := &common.UserWithID{
		Id:      id.String(),
		Email:   user.Email,
		Name:    user.Name,
		Surname: user.Surname,
	}
	s.data = append(s.data, newUser)
	return nil
}

func (s *InMemoryUserStorage) UpdateUser(userID string, user common.UserWithoutID) bool {
	for _, u := range s.data {
		if u.Id == userID {
			u.Name = user.Name
			u.Surname = user.Surname
			u.Email = user.Email
			return true
		}
	}

	return false
}

func (s *InMemoryUserStorage) DeleteUser(userID string) bool {
	for index, u := range s.data {
		if u.Id == userID {
			copy(s.data[index:], s.data[index+1:])
			s.data = s.data[:len(s.data)-1]
			return true
		}
	}

	return false
}

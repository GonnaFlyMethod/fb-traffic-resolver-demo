package server

import "github.com/google/uuid"

type InMemoryUserStorage struct {
	data []map[string]string
}

func (s *InMemoryUserStorage) GetAllUsers() []map[string]string {
	return s.data
}

func (s *InMemoryUserStorage) CreateNewUser(userData User) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	newUser := map[string]string{
		"id":              id.String(),
		"email":           userData.Email,
		"name":            userData.Name,
		"surname":         userData.Surname,
		"password":        userData.Password,
		"confirmPassword": userData.ConfirmPassword,
		"avatar":          userData.Avatar,
	}
	s.data = append(s.data, newUser)

	return nil
}

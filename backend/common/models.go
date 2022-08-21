package common

type UserWithID struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type UserWithoutID struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (u *UserWithoutID) IsCompletelyFilled() bool {
	if u.Email != "" && u.Name != "" && u.Surname != "" {
		return true
	}

	return false
}

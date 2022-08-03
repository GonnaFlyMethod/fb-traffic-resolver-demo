package server

type User struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Avatar          string `json:"avatar"`
}

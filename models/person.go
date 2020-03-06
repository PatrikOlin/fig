package models

type Person struct {
	Name     string `json:"name"`
	Pin      string `json:"pin"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

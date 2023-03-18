package userinput

import "time"

type CreateUserInput struct {
	User CreateUser `json:"mjUser"`
}

type CreateUser struct {
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Gender   string    `json:"gender"`
	Address  string    `json:"address"`
	Email    string    `json:"email"`
}

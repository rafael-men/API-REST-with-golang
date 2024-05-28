package types

import "time"

type User struct {
	ID        int       `json:"id"`
	Firstname string    `json:"Firstname"`
	Lastname  string    `json:"Lastname"`
	Email     string    `json:"Email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdat"`
}

type RegisterUserPayload struct {
	Firstname string `json:"Firstname`
	Lastname  string `json:"Lastname`
	Email     string `json:"Email`
	Password  string `json:"Password`
}

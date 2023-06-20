package entities

import (
	"time"
)

type User struct {
	Id        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Name      string `json:"name,omitempty"`
	BirthDate string `json:"birth_date,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// NewUser return an instance of User
func NewUser(email string, name string, birthDate string) *User {
	return &User{
		Email:     email,
		Name:      name,
		BirthDate: birthDate,
	}
}

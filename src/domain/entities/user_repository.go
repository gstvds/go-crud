package entities

import (
	"context"
)

type UserRepository interface {
	Create(*User, context.Context) (*User, error)
	GetByEmail(string, context.Context) (*User, error)
	GetById(string, context.Context) (*User, error)
	Update(*User, context.Context) (*User, error)
	Delete(string, context.Context) error
}

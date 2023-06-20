package repositories

import (
	"context"
	"go-crud/src/domain/entities"
)

type UserRepository interface {
	Create(*entities.User, context.Context) (*entities.User, error)
	GetByEmail(*entities.User, context.Context) (*entities.User, error)
	GetById(*entities.User, context.Context) (*entities.User, error)
	Update(*entities.User, context.Context) (*entities.User, error)
	Delete(string, context.Context) error
}

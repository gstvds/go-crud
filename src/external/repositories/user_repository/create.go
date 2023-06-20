package userrepository

import (
	"context"
	"go-crud/prisma/db"
	"go-crud/src/domain/entities"
	"go-crud/src/external/repositories/mappers"
)

func (repository repository) Create(user *entities.User, ctx context.Context) (*entities.User, error) {
	if createdUser, err := repository.db.User.CreateOne(
		db.User.Email.Set(user.Email),
		db.User.Name.Set(user.Name),
		db.User.BirthDate.Set(user.BirthDate),
	).Exec(ctx); err != nil {
		return nil, err
	} else {
		return mappers.ToDomain(createdUser), nil
	}
}

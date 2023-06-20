package userrepository

import (
	"context"
	"go-crud/prisma/db"
	"go-crud/src/domain/entities"
	"go-crud/src/external/repositories/mappers"
)

func (repository repository) Update(user *entities.User, ctx context.Context) (*entities.User, error) {
	if updatedUser, err := repository.db.User.FindUnique(
		db.User.ID.Equals(user.Id),
	).Update(
		db.User.Email.Set(user.Email),
		db.User.Name.Set(user.Name),
		db.User.BirthDate.Set(user.BirthDate),
	).Exec(ctx); err != nil {
		return nil, err
	} else {
		return mappers.ToDomain(updatedUser), nil
	}
}

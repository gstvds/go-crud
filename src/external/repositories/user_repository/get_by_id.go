package userrepository

import (
	"context"
	"go-crud/prisma/db"
	"go-crud/src/domain/entities"
	"go-crud/src/external/repositories/mappers"
)

func (repository repository) GetById(user *entities.User, ctx context.Context) (*entities.User, error) {
	if user, err := repository.db.User.FindUnique(db.User.ID.Equals(user.Id)).Exec(ctx); err != nil {
		return nil, err
	} else {
		return mappers.ToDomain(user), err
	}
}

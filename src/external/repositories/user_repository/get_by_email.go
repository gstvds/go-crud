package userrepository

import (
	"context"
	"go-crud/prisma/db"
	"go-crud/src/domain/entities"
	"go-crud/src/external/repositories/mappers"
)

func (repository repository) GetByEmail(user *entities.User, ctx context.Context) (*entities.User, error) {
	if foundUser, err := repository.db.User.FindFirst(db.User.Email.Equals(user.Email)).Exec(ctx); err != nil {
		return nil, err
	} else {
		return mappers.ToDomain(foundUser), nil
	}
}

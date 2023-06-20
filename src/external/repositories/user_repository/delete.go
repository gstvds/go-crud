package userrepository

import (
	"context"
	"go-crud/prisma/db"
)

func (repository repository) Delete(userId string, ctx context.Context) error {
	_, err := repository.db.User.FindUnique(
		db.User.ID.Equals(userId),
	).Delete().Exec(ctx)
	return err
}

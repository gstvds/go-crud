package mappers

import (
	"go-crud/prisma/db"
	"go-crud/src/domain/entities"
)

func ToDomain(dbUser *db.UserModel) *entities.User {
	var user = entities.User{
		Id:        dbUser.ID,
		Email:     dbUser.Email,
		Name:      dbUser.Name,
		BirthDate: dbUser.BirthDate,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}

	return &user
}

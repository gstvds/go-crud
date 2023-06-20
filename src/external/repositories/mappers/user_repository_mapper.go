package mappers

import (
	"go-crud/prisma/db"
	"go-crud/src/domain/entities"
)

type UserRepositoryMapper struct{}

func NewUserRepositoryMapper() *UserRepositoryMapper {
	return &UserRepositoryMapper{}
}

func (UserRepositoryMapper) ToDomain(dbUser *db.UserModel) *entities.User {
	return &entities.User{
		Id:        dbUser.ID,
		Email:     dbUser.Email,
		Name:      dbUser.Name,
		BirthDate: dbUser.BirthDate,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}

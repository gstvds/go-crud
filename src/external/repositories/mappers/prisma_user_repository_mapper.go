package mappers

import (
	"go-crud/prisma/db"
	"go-crud/src/domain/entities"
)

type PrismaUserRepositoryMapper struct{}

func NewPrismaUserRepositoryMapper() *PrismaUserRepositoryMapper {
	return &PrismaUserRepositoryMapper{}
}

func (PrismaUserRepositoryMapper) ToDomain(dbUser *db.UserModel) *entities.User {
	return &entities.User{
		Id:        dbUser.ID,
		Email:     dbUser.Email,
		Name:      dbUser.Name,
		BirthDate: dbUser.BirthDate,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}

package mappers

import (
	"go-crud/prisma/db"
	"go-crud/src/domain/entities"
)

type PrismaContactRepositoryMapper struct{}

func NewPrismaContactRepositoryMapper() *PrismaContactRepositoryMapper {
	return &PrismaContactRepositoryMapper{}
}

func (PrismaContactRepositoryMapper) ToDomain(dbContact *db.ContactModel) *entities.Contact {
	userRepositoryMapper := NewPrismaUserRepositoryMapper()

	return &entities.Contact{
		Id:        dbContact.ID,
		Channel:   entities.CHANNEL(dbContact.Channel),
		Enabled:   dbContact.Enabled,
		Receiver:  dbContact.Receiver,
		UserId:    dbContact.UserID,
		User:      *userRepositoryMapper.ToDomain(dbContact.User()),
		CreatedAt: dbContact.CreatedAt,
		UpdatedAt: dbContact.UpdatedAt,
	}
}

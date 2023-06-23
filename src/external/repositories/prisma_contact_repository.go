package repositories

import (
	"context"
	"go-crud/prisma/db"
	"go-crud/src/domain/entities"
	"go-crud/src/external/repositories/mappers"
)

type PrismaContactRepository struct {
	db *db.PrismaClient
}

func NewPrismaContactRepository(db *db.PrismaClient) *PrismaContactRepository {
	return &PrismaContactRepository{db}
}

func (contactRepository PrismaContactRepository) Create(contact *entities.Contact, ctx context.Context) (*entities.Contact, error) {
	if createdContact, err := contactRepository.db.Contact.CreateOne(
		db.Contact.Channel.Set(db.CHANNEL(contact.Channel)),
		db.Contact.Receiver.Set(contact.Receiver),
		db.Contact.User.Link(
			db.User.ID.Equals(contact.UserId),
		),
	).With(db.Contact.User.Fetch()).Exec(ctx); err != nil {
		return nil, err
	} else {
		contactRepositoryMapper := mappers.NewPrismaContactRepositoryMapper()
		return contactRepositoryMapper.ToDomain(createdContact), nil
	}
}

func (contactRepository PrismaContactRepository) GetByUserId(userId string, ctx context.Context) (*[]entities.Contact, error) {
	if foundContacts, err := contactRepository.db.Contact.FindMany(
		db.Contact.UserID.Equals(userId),
	).With(db.Contact.User.Fetch()).Exec(ctx); err != nil {
		return nil, err
	} else {
		prismaContactRepositoryMapper := mappers.NewPrismaContactRepositoryMapper()
		var contacts []entities.Contact
		for _, contact := range foundContacts {
			contacts = append(contacts, *prismaContactRepositoryMapper.ToDomain(&contact))
		}

		return &contacts, nil
	}
}

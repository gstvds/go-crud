package usecases

import (
	"context"
	"fmt"
	"go-crud/src/domain/entities"
	"time"
)

type CreateContactInputDTO struct {
	Email  string
	UserId string
}

type CreateContactOutputDTO struct {
	Id        string           `json:"id,omitempty"`
	Channel   entities.CHANNEL `json:"channel,omitempty"`
	Enabled   bool             `json:"enabled,omitempty"`
	Receiver  string           `json:"receiver,omitempty"`
	UserId    string           `json:"user_id,omitempty"`
	User      entities.User    `json:"user,omitempty"`
	CreatedAt time.Time        `json:"created_at,omitempty"`
	UpdatedAt time.Time        `json:"updated_at,omitempty"`
}

type CreateContactUseCase struct {
	ContactRepository entities.ContactRepository
}

func NewCreateContactUseCase(contactRepository entities.ContactRepository) *CreateContactUseCase {
	return &CreateContactUseCase{ContactRepository: contactRepository}
}

func (createContactUseCase CreateContactUseCase) Exec(input CreateContactInputDTO) (*[]CreateContactOutputDTO, error) {
	ctx := context.Background()

	emailContact := entities.NewContact(entities.CHANNEL(entities.EMAIL), true, input.Email, input.UserId)
	pushContact := entities.NewContact(entities.CHANNEL(entities.PUSH_NOTIFICATION), true, "fake_token", input.UserId)

	var output []CreateContactOutputDTO

	if createContact, err := createContactUseCase.ContactRepository.Create(emailContact, ctx); err != nil {
		fmt.Println(createContact)
		output = append(output, CreateContactOutputDTO{
			Id:        createContact.Id,
			Channel:   createContact.Channel,
			Enabled:   createContact.Enabled,
			Receiver:  createContact.Receiver,
			UserId:    createContact.UserId,
			CreatedAt: createContact.CreatedAt,
			UpdatedAt: createContact.UpdatedAt,
		})
	}
	if createContact, err := createContactUseCase.ContactRepository.Create(pushContact, ctx); err != nil {
		output = append(output, CreateContactOutputDTO{
			Id:        createContact.Id,
			Channel:   createContact.Channel,
			Enabled:   createContact.Enabled,
			Receiver:  createContact.Receiver,
			UserId:    createContact.UserId,
			CreatedAt: createContact.CreatedAt,
			UpdatedAt: createContact.UpdatedAt,
		})
	}

	return &output, nil
}

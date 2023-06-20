package usecases

import (
	"context"
	"fmt"
	"go-crud/src/domain/entities"
	"time"
)

type UpsertUserInputDTO struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
}

type UpsertUserOutputDTO struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	BirthDate string    `json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpsertUserUseCase struct {
	UserRepository entities.UserRepository
}

func NewUpsertUserUseCase(userRepository entities.UserRepository) *UpsertUserUseCase {
	return &UpsertUserUseCase{UserRepository: userRepository}
}

// UpsertUserUseCase updates a user (or creates if doesn't exists)
func (upsertUserUseCase UpsertUserUseCase) Exec(input UpsertUserInputDTO) (*UpsertUserOutputDTO, error) {
	user := entities.NewUser(input.Email, input.Name, input.BirthDate)
	user.Id = input.Id
	ctx := context.Background()

	foundUser, _ := upsertUserUseCase.UserRepository.GetById(user.Id, ctx)
	if foundUser == nil {
		// If doesn't exist, create a new user
		if createdUser, err := upsertUserUseCase.UserRepository.Create(user, ctx); err != nil {
			return nil, err
		} else {
			fmt.Println("User created successfully")
			return &UpsertUserOutputDTO{
				Id:        createdUser.Id,
				Name:      createdUser.Name,
				Email:     createdUser.Email,
				BirthDate: createdUser.BirthDate,
				CreatedAt: createdUser.CreatedAt,
				UpdatedAt: createdUser.UpdatedAt,
			}, nil
		}
	} else {
		if updatedUser, err := upsertUserUseCase.UserRepository.Update(user, ctx); err != nil {
			return nil, err
		} else {
			fmt.Println("User updated successfully")
			return &UpsertUserOutputDTO{
				Id:        updatedUser.Id,
				Name:      updatedUser.Name,
				Email:     updatedUser.Email,
				BirthDate: updatedUser.BirthDate,
				CreatedAt: updatedUser.CreatedAt,
				UpdatedAt: updatedUser.UpdatedAt,
			}, nil
		}
	}
}

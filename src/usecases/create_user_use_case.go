package usecases

import (
	"context"
	"go-crud/src/domain/entities"
	"go-crud/src/domain/repositories"
	"log"
	"time"
)

type CreateUserInputDTO struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
}

type CreateUserOutputDTO struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	BirthDate string    `json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserUseCase struct {
	UserRepository repositories.UserRepository
}

func NewCreateUserUseCase(userRepository repositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{UserRepository: userRepository}
}

// Exec creates a new User
func (useCase CreateUserUseCase) Exec(input CreateUserInputDTO) (*CreateUserOutputDTO, error) {
	user := entities.NewUser(input.Email, input.Name, input.BirthDate)
	ctx := context.Background()

	foundUser, _ := useCase.UserRepository.GetByEmail(user, ctx)
	if foundUser == nil {
		log.Println("User not found. Creating one...")
		createdUser, err := useCase.UserRepository.Create(user, ctx)
		if err != nil {
			return nil, err
		}

		return &CreateUserOutputDTO{
			Id:        createdUser.Id,
			Email:     createdUser.Email,
			Name:      createdUser.Name,
			BirthDate: createdUser.BirthDate,
			CreatedAt: createdUser.CreatedAt,
			UpdatedAt: createdUser.UpdatedAt,
		}, nil
	}

	return &CreateUserOutputDTO{
		Id:        foundUser.Id,
		Email:     foundUser.Email,
		Name:      foundUser.Name,
		BirthDate: foundUser.BirthDate,
		CreatedAt: foundUser.CreatedAt,
		UpdatedAt: foundUser.UpdatedAt,
	}, nil
}

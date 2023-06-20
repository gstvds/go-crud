package usecases

import (
	"context"
	"go-crud/src/domain/entities"
	"time"
)

type GetUserInputDTO struct {
	UserId string `json:"user_id"`
}

type GetUserOutputDTO struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	BirthDate string    `json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserUseCase struct {
	UserRepository entities.UserRepository
}

func NewGetUserUseCase(userRepository entities.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{UserRepository: userRepository}
}

// GetUserUseCase returns a user by its email
func (getUserUseCase GetUserUseCase) Exec(input GetUserInputDTO) (*GetUserOutputDTO, error) {
	userId := input.UserId
	ctx := context.Background()

	if foundUser, err := getUserUseCase.UserRepository.GetById(userId, ctx); err != nil {
		return nil, err
	} else {
		return &GetUserOutputDTO{
			Id:        foundUser.Id,
			Name:      foundUser.Name,
			Email:     foundUser.Email,
			BirthDate: foundUser.BirthDate,
			CreatedAt: foundUser.CreatedAt,
			UpdatedAt: foundUser.UpdatedAt,
		}, nil
	}

}

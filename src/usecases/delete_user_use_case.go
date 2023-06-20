package usecases

import (
	"context"
	"fmt"
	"go-crud/src/domain/entities"
)

type DeleteUserInputDTO struct {
	UserId string `json:"user_id"`
}

type DeleteUserUseCase struct {
	UserRepository entities.UserRepository
}

func NewDeleteUserUseCase(userRepository entities.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{UserRepository: userRepository}
}

// Exec the DeleteUserUseCase to delete a user by its email
func (deleteUserUseCase DeleteUserUseCase) Exec(input DeleteUserInputDTO) error {
	userId := input.UserId
	ctx := context.Background()

	if _, err := deleteUserUseCase.UserRepository.GetById(userId, ctx); err != nil {
		if err.Error() == "ErrNotFound" {
			fmt.Println("User not found")
			return err
		}
	}

	if err := deleteUserUseCase.UserRepository.Delete(userId, ctx); err != nil {
		return err
	}

	fmt.Println("User deleted successfully")
	return nil
}

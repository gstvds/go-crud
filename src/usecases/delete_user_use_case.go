package usecases

import (
	"context"
	"errors"
	"fmt"
	"go-crud/src/domain/entities"
	"go-crud/src/external/providers/database"
	"go-crud/src/external/repositories"

	"gorm.io/gorm"
)

// DeleteUserUseCase deletes a user by its email
func DeleteUserUseCase(user *entities.User) error {
	db := database.Get()
	userRepository := repositories.NewPrismaUserRepository(db)
	ctx := context.Background()

	existingUser := user

	if _, err := userRepository.GetById(existingUser, ctx); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User not found")
			return err
		}
	}

	if err := userRepository.Delete(user.Id, ctx); err != nil {
		return err
	}

	fmt.Println("User deleted successfully")
	return nil
}

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

// UpsertUserUseCase updates a user (or creates if doesn't exists)
func UpsertUserUseCase(user *entities.User) error {
	db := database.Get()
	userRepository := repositories.NewPrismaUserRepository(db)
	ctx := context.Background()

	existingUser := *user

	if _, err := userRepository.GetById(&existingUser, ctx); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If doesn't exist, create a new user
			if _, err = userRepository.Create(user, ctx); err != nil {
				return err
			}

			fmt.Println("User created successfully")
			return nil
		}
		return err
	}

	if _, err := userRepository.Update(user, ctx); err != nil {
		return err
	}

	fmt.Println("User updated successfully")

	return nil
}

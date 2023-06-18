package usecases

import (
	"errors"
	"fmt"
	"go-crud/src/domain/entities"
	"go-crud/src/external/providers/database"
	userrepository "go-crud/src/external/repositories/user_repository"

	"gorm.io/gorm"
)

// UpsertUserUseCase updates a user (or creates if doesn't exists)
func UpsertUserUseCase(user *entities.User) error {
	db := database.Get()
	userRepository := userrepository.New(db)

	existingUser := *user

	if err := userRepository.GetById(&existingUser); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If doesn't exist, create a new user
			if err = userRepository.Create(user); err != nil {
				return err
			}

			fmt.Println("User created successfully")
			return nil
		}
		return err
	}

	if err := userRepository.Update(user); err != nil {
		return err
	}

	fmt.Println("User updated successfully")

	return nil
}

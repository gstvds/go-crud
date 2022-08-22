package services

import (
	"errors"
	"fmt"
	"go-crud/src/infra/repositories/models"
	userrepository "go-crud/src/infra/repositories/user_repository"
	"go-crud/src/providers/database"

	"gorm.io/gorm"
)

// UpsertUserService updates a user (or creates if doesn't exists)
func UpsertUserService(user *models.User) error {
	db := database.Get()
	userRepository := userrepository.New(db)

	existingUser := *user

	if err := userRepository.Read(&existingUser); err != nil {
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

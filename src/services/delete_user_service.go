package services

import (
	"errors"
	"fmt"
	"go-crud/src/infra/repositories/models"
	userrepository "go-crud/src/infra/repositories/user_repository"
	"go-crud/src/providers/database"

	"gorm.io/gorm"
)

// DeleteUserService deletes a user by its email
func DeleteUserService(user *models.User) error {
	db := database.Get()
	userRepository := userrepository.New(db)

	existingUser := user

	if err := userRepository.Read(existingUser); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User not found")
			return err
		}
	}

	if err := userRepository.Delete(user); err != nil {
		return err
	}

	fmt.Println("User deleted successfully")
	return nil
}

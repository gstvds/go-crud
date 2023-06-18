package usecases

import (
	"errors"
	"fmt"
	"go-crud/src/domain/entities"
	"go-crud/src/external/providers/database"
	userrepository "go-crud/src/external/repositories/user_repository"

	"gorm.io/gorm"
)

// DeleteUserUseCase deletes a user by its email
func DeleteUserUseCase(user *entities.User) error {
	db := database.Get()
	userRepository := userrepository.New(db)

	existingUser := user

	if err := userRepository.GetById(existingUser); err != nil {
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

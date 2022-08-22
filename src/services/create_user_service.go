package services

import (
	"errors"
	"go-crud/src/infra/repositories/models"
	userrepository "go-crud/src/infra/repositories/user_repository"
	"go-crud/src/providers/database"
	"log"

	"gorm.io/gorm"
)

// CreateUserService creates a new user
func CreateUserService(user *models.User) error {
	db := database.Get()
	userRepository := userrepository.New(db)

	existingUser := user

	// Check if user already exists before attempting to create
	if err := userRepository.Read(existingUser); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("User not found. Creating one...")
			if err = userRepository.Create(user); err != nil {
				return err
			}
		} else {
			log.Println("Something went wrong")
			return err
		}
	}

	return nil
}

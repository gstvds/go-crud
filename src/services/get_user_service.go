package services

import (
	"go-crud/src/infra/repositories/models"
	userrepository "go-crud/src/infra/repositories/user_repository"
	"go-crud/src/providers/database"
)

// GetUserService returns a user by its email
func GetUserService(user *models.User) error {
	db := database.Get()
	userRepository := userrepository.New(db)

	if err := userRepository.Read(user); err != nil {
		return err
	}

	return nil
}

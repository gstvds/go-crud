package usecases

import (
	"go-crud/src/domain/entities"
	"go-crud/src/external/providers/database"
	userrepository "go-crud/src/external/repositories/user_repository"
)

// GetUserUseCase returns a user by its email
func GetUserUseCase(user *entities.User) error {
	db := database.Get()
	userRepository := userrepository.New(db)

	if err := userRepository.GetById(user); err != nil {
		return err
	}

	return nil
}

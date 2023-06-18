package usecases

import (
	"errors"
	"go-crud/src/domain/entities"
	"go-crud/src/external/providers/database"
	userrepository "go-crud/src/external/repositories/user_repository"
	"log"

	"gorm.io/gorm"
)

// CreateUserUseCase creates a new user
func CreateUserUseCase(user *entities.User) error {
	db := database.Get()
	userRepository := userrepository.New(db)

	existingUser := user

	if err := userRepository.GetByEmail(existingUser); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("User not found. Creating one...")
			if err := userRepository.Create(user); err != nil {
				return err
			}
		}
	}

	return errors.New("user already exists")
}

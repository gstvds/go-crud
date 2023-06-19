package usecases

import (
	"context"
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
	ctx := context.Background()

	existingUser := user

	if _, err := userRepository.GetByEmail(existingUser, ctx); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("User not found. Creating one...")
			if err := userRepository.Create(user); err != nil {
				return err
			}
		}
	}

	return errors.New("user already exists")
}

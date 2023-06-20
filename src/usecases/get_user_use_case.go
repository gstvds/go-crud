package usecases

import (
	"context"
	"go-crud/src/domain/entities"
	"go-crud/src/external/providers/database"
	"go-crud/src/external/repositories"
)

// GetUserUseCase returns a user by its email
func GetUserUseCase(user *entities.User) error {
	db := database.Get()
	userRepository := repositories.NewPrismaUserRepository(db)
	ctx := context.Background()

	if _, err := userRepository.GetById(user, ctx); err != nil {
		return err
	}

	return nil
}

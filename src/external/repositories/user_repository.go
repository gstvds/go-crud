package repositories

import (
	"context"
	"go-crud/prisma/db"
	"go-crud/src/domain/entities"
	"go-crud/src/external/repositories/mappers"
)

type PrismaUserRepository struct {
	db *db.PrismaClient
}

// NewPrismaUserRepository create a new PrismaUserRepository instance
func NewPrismaUserRepository(db *db.PrismaClient) *PrismaUserRepository {
	return &PrismaUserRepository{db}
}

func (userRepository PrismaUserRepository) Create(user *entities.User, ctx context.Context) (*entities.User, error) {
	if createdUser, err := userRepository.db.User.CreateOne(
		db.User.Email.Set(user.Email),
		db.User.Name.Set(user.Name),
		db.User.BirthDate.Set(user.BirthDate),
	).Exec(ctx); err != nil {
		return nil, err
	} else {
		userRepositoryMapper := mappers.NewUserRepositoryMapper()
		return userRepositoryMapper.ToDomain(createdUser), nil
	}
}

func (repository PrismaUserRepository) Delete(userId string, ctx context.Context) error {
	_, err := repository.db.User.FindUnique(
		db.User.ID.Equals(userId),
	).Delete().Exec(ctx)
	return err
}

func (repository PrismaUserRepository) GetByEmail(user *entities.User, ctx context.Context) (*entities.User, error) {
	if foundUser, err := repository.db.User.FindFirst(db.User.Email.Equals(user.Email)).Exec(ctx); err != nil {
		return nil, err
	} else {
		userRepositoryMapper := mappers.NewUserRepositoryMapper()
		return userRepositoryMapper.ToDomain(foundUser), nil
	}
}

func (repository PrismaUserRepository) GetById(user *entities.User, ctx context.Context) (*entities.User, error) {
	if user, err := repository.db.User.FindUnique(db.User.ID.Equals(user.Id)).Exec(ctx); err != nil {
		return nil, err
	} else {
		userRepositoryMapper := mappers.NewUserRepositoryMapper()
		return userRepositoryMapper.ToDomain(user), err
	}
}

func (repository PrismaUserRepository) Update(user *entities.User, ctx context.Context) (*entities.User, error) {
	if updatedUser, err := repository.db.User.FindUnique(
		db.User.ID.Equals(user.Id),
	).Update(
		db.User.Email.Set(user.Email),
		db.User.Name.Set(user.Name),
		db.User.BirthDate.Set(user.BirthDate),
	).Exec(ctx); err != nil {
		return nil, err
	} else {
		userRepositoryMapper := mappers.NewUserRepositoryMapper()
		return userRepositoryMapper.ToDomain(updatedUser), nil
	}
}

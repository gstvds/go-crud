package repositories

import "go-crud/src/domain/entities"

type UserRepository interface {
	Create(user *entities.User) error
	GetByEmail(user *entities.User) error
	GetById(user *entities.User) error
	Update(user *entities.User) error
	Delete(id string) error
}

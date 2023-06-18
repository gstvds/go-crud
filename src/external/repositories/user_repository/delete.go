package userrepository

import "go-crud/src/domain/entities"

func (repository repository) Delete(user *entities.User) error {
	return repository.db.Delete(user).Error
}

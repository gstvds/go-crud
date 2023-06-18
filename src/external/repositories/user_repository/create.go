package userrepository

import "go-crud/src/domain/entities"

func (repository repository) Create(user *entities.User) error {
	if err := repository.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

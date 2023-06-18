package userrepository

import (
	"go-crud/src/domain/entities"
)

func (repository repository) GetByEmail(user *entities.User) error {
	if err := repository.db.Model(user).Where("email = ?", user.Email).First(user).Error; err != nil {
		return err
	}

	return nil
}

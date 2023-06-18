package userrepository

import "go-crud/src/domain/entities"

func (repository repository) GetById(user *entities.User) error {
	if err := repository.db.Where("id = ?", user.Id).First(user).Error; err != nil {
		return err
	} else {
		return nil
	}
}

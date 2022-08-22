package userrepository

import (
	"go-crud/src/infra/repositories/models"

	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

// New repository instance
func New(db *gorm.DB) *User {
	return &User{db}
}

// Create a new user
func (repository User) Create(user *models.User) error {
	return repository.db.Create(user).Error
}

// Read a user record
func (repository User) Read(user *models.User) error {
	return repository.db.Where("email = ?", user.Email).First(user).Error
}

// Update a user
func (repository User) Update(user *models.User) error {
	return repository.db.Model(user).Where("email = ?", user.Email).Updates(user).Error
}

// Delete a user
func (repository User) Delete(user *models.User) error {
	return repository.db.Delete(user).Error
}

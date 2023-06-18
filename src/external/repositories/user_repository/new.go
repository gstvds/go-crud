package userrepository

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// New repository instance
func New(db *gorm.DB) *repository {
	return &repository{db}
}

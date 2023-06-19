package userrepository

import (
	"go-crud/prisma/db"
)

type repository struct {
	db *db.PrismaClient
}

// New repository instance
func New(db *db.PrismaClient) *repository {
	return &repository{db}
}

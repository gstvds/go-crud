package database

import (
	"go-crud/prisma/db"
	"log"
)

var client *db.PrismaClient

func SetupDatabase() {
	client = db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Fatal(err)
	}
}

func Get() *db.PrismaClient {
	return client
}

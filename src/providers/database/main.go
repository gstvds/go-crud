package database

import (
	"fmt"
	"go-crud/src/infra/repositories/models"
	"go-crud/src/utils"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// SetupDatabase prepare the database connection
func SetupDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		utils.DATABASE.Host,
		utils.DATABASE.User,
		utils.DATABASE.Password,
		utils.DATABASE.Name,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}))

	if err != nil {
		log.Fatal(err)
	}

	dbConfig, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	dbConfig.SetMaxOpenConns(1000)
	dbConfig.SetMaxIdleConns(50)
	dbConfig.SetConnMaxLifetime(5 * time.Minute)

	if err = db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
	}

	DB = db
}

// Get the created database instance
func Get() *gorm.DB {
	return DB
}

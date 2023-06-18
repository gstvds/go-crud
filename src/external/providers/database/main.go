package database

import (
	"fmt"
	"go-crud/src/domain/entities"
	"go-crud/src/shared"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// SetupDatabase prepare the database connection
func SetupDatabase() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		shared.DATABASE.Host,
		shared.DATABASE.Port,
		shared.DATABASE.User,
		shared.DATABASE.Password,
		shared.DATABASE.Name,
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

	if err = db.AutoMigrate(&entities.User{}); err != nil {
		log.Fatal(err)
	}

	DB = db
}

// Get the created database instance
func Get() *gorm.DB {
	return DB
}

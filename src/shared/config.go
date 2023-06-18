package shared

import (
	"os"

	"github.com/joho/godotenv"
)

type databaseEnv struct {
	User     string
	Port     string
	Password string
	Name     string
	Host     string
}

var (
	DATABASE databaseEnv
)

// LoadEnv loads all environment variables
func LoadEnv() {
	godotenv.Load()

	DATABASE.User = os.Getenv("PSQL_USERNAME")
	DATABASE.Port = os.Getenv("PSQL_PORT")
	DATABASE.Password = os.Getenv("PSQL_PASSWORD")
	DATABASE.Name = os.Getenv("PSQL_DATABASE_NAME")
	DATABASE.Host = os.Getenv("PSQL_HOST")
}

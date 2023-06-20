package config

import (
	"os"

	"github.com/joho/godotenv"
)

type environmentVariables struct {
	PORT                    string
	KAFKA_BOOTSTRAP_SERVERS string
	KAFKA_PROTOCOL          string
	KAFKA_USERNAME          string
	KAFKA_PASSWORD          string
}

var (
	ENVRIONMENT_VARIABLES environmentVariables
)

// LoadEnv loads all environment variables
func LoadEnv() {
	godotenv.Load()

	ENVRIONMENT_VARIABLES.PORT = os.Getenv("PORT")
	ENVRIONMENT_VARIABLES.KAFKA_BOOTSTRAP_SERVERS = os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	ENVRIONMENT_VARIABLES.KAFKA_PROTOCOL = os.Getenv("KAFKA_PROTOCOL")
	ENVRIONMENT_VARIABLES.KAFKA_USERNAME = os.Getenv("KAFKA_USERNAME")
	ENVRIONMENT_VARIABLES.KAFKA_PASSWORD = os.Getenv("KAFKA_PASSWORD")
}

package main

import (
	"fmt"
	"go-crud/src/external/providers/database"
	"go-crud/src/external/providers/routes"
	"go-crud/src/shared"
	"log"
	"os"
)

// configureRouter configures server port and starts the server
func configureRouter() {
	app := routes.NewRouter()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3333"
	}

	log.Printf("Server started on PORT: %s", PORT)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", PORT)))
}

func main() {
	shared.LoadEnv()
	database.SetupDatabase()

	configureRouter()
}

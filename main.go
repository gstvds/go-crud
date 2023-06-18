package main

import (
	"fmt"
	"go-crud/src/external/providers/database"
	"go-crud/src/external/providers/routes"
	"go-crud/src/shared"
	"log"
	"net/http"
	"os"
)

// configureRouter configures server port and starts the server
func configureRouter() {
	serverRouter := routes.Generate()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3333"
	}

	log.Printf("Server started on PORT: %s", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), serverRouter))
}

func main() {
	shared.LoadEnv()
	database.SetupDatabase()

	configureRouter()
}

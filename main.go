package main

import (
	"fmt"
	"go-crud/src/infra/http/routes"
	"go-crud/src/providers/database"
	"go-crud/src/utils"
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
	utils.LoadEnv()
	database.SetupDatabase()

	configureRouter()
}

package main

import (
	"docker-project-api/cmd/api"
	"docker-project-api/internal/db"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dbInstance := db.Connect()
	api.CreateServer()

	defer dbInstance.Close()
}

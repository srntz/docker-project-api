package main

import (
	"docker-project-api/cmd/api"
	"docker-project-api/internal/db"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dbInstance := db.CreateConnection()
	api.CreateServer(dbInstance)

	defer dbInstance.Close()
}

package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

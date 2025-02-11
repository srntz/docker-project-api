package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var connection *sql.DB

func Connect() *sql.DB {
	if connection == nil {
		db, err := sql.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))

		if err != nil {
			log.Fatal(err)
		}

		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}

		connection = db
	}

	return connection
}

package api

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateServer(db *sql.DB) {
	if db != nil {
		fmt.Println("Server created.")
	} else {
		log.Fatal("Server could not connect to the database.")
	}
}

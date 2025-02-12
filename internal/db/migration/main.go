package main

import (
	"database/sql"
	"docker-project-api/internal/db"
	"docker-project-api/internal/models"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

// File-scoped database instance. The variable gets initialized from main()
var dbInstance *sql.DB

// Mock data for the student table
var studentData = []models.StudentInsert{
	{"ABC10000", "John Doe", "CPRG-101"},
	{"ABC10001", "Jane Doe", "CPRG-102"},
}

/*
Drops all tables in the database
*/
func dropAll() {
	query := "DROP TABLE IF EXISTS student CASCADE;"

	_, err := dbInstance.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tables dropped.")
}

/*
Creates tables
*/
func loadSchema() {
	query := `CREATE TABLE IF NOT EXISTS student (
		    student_id VARCHAR(128) PRIMARY KEY,
		    student_name VARCHAR(255),
		    course VARCHAR(255),
		    present_date TIMESTAMP DEFAULT NOW())`

	_, err := dbInstance.Exec(query)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Tables created.")
}

/*
Loads mock data into the tables
*/
func loadData() {
	query := `INSERT INTO student (student_id, student_name, course) VALUES ($1, $2, $3) RETURNING student_id`
	for i := 0; i < len(studentData); i++ {
		var id string
		err := dbInstance.QueryRow(query, studentData[i].StudentId, studentData[i].StudentName, studentData[i].Course).Scan(&id)

		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Printf("A student with id: %s created.\n", id)
	}

	fmt.Println("All data has been loaded successfully.")
}

func main() {
	dbInstance = db.Connect()
	dropAll()
	loadSchema()
	loadData()
}

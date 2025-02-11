package api

import (
	"docker-project-api/cmd/api/handlers"
	"log"
	"net/http"
)

func CreateServer() {
	http.HandleFunc("/api/v1/getStudents", handlers.StudentHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

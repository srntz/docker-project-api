package api

import (
	"docker-project-api/cmd/api/handlers"
	"log"
	"net/http"
)

func CreateServer() {
	http.HandleFunc("/api/v1/getAllStudents", handlers.GetAllStudents)
	http.HandleFunc("/api/v1/getStudent", handlers.GetStudent)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

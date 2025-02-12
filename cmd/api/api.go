package api

import (
	"docker-project-api/cmd/api/handlers/student"
	"net/http"
)

func CreateServer() {
	router := http.NewServeMux()

	router.HandleFunc("GET /getAllStudents", student.GetAllStudents)
	router.HandleFunc("GET /getStudent", student.GetStudent)
	router.HandleFunc("POST /student", student.CreateStudent)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

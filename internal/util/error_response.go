package util

import (
	"docker-project-api/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, error error, statusCode int, message string) {
	log.Println(error)
	
	response := models.ErrorResponse{Error: struct {
		StatusCode int    `json:"status_code"`
		Message    string `json:"message"`
	}{StatusCode: statusCode, Message: message}}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

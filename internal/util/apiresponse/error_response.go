package apiresponse

import (
	"docker-project-api/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

const MessageInvalidQuery = "Invalid database query"
const MessageScanError = "Data mapping failed"
const MessageResponseEncodingError = "Response encoding failed"
const MessageInternalError = "An internal server error occurred"

func SendErrorResponse(w http.ResponseWriter, error error, statusCode int, message string) {
	if error != nil {
		log.Println(error)
	}

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

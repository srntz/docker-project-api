package handlers

import (
	"fmt"
	"net/http"
)

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Student handler")
}

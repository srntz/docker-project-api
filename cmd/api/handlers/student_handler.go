package handlers

import (
	"docker-project-api/internal/db"
	"docker-project-api/internal/models"
	"docker-project-api/internal/util"
	"encoding/json"
	"net/http"
)

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	dbInstance := db.Connect()
	response := models.Response[models.Student]{[]models.Student{}}

	query := "SELECT student_id, student_name, course, present_date FROM student"

	rows, err := dbInstance.Query(query)
	if err != nil {
		util.SendErrorResponse(w, err, http.StatusInternalServerError, "Invalid database query")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.StudentId, &student.StudentName, &student.Course, &student.PresentDate)
		if err != nil {
			util.SendErrorResponse(w, err, http.StatusInternalServerError, "Data could not be mapped to a model")
		}

		response.Data = append(response.Data, student)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		util.SendErrorResponse(w, err, http.StatusInternalServerError, "Response encoding failed")
	}
}

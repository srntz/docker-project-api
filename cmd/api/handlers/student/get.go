package student

import (
	"database/sql"
	"docker-project-api/internal/db"
	"docker-project-api/internal/models"
	"docker-project-api/internal/util/apiresponse"
	"encoding/json"
	"errors"
	"net/http"
)

func GetStudent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		apiresponse.SendErrorResponse(w, nil, http.StatusBadRequest, "Student ID is required. Use a URL parameter ('id')")
		return
	}

	dbInstance := db.Connect()

	query := "SELECT student_id, student_name, course, present_date FROM student WHERE student_id = $1"

	response := models.Response[models.Student]{}
	student := models.Student{}
	err := dbInstance.QueryRow(query, id).Scan(&student.StudentId, &student.StudentName, &student.Course, &student.PresentDate)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		apiresponse.SendErrorResponse(w, err, http.StatusInternalServerError, apiresponse.MessageScanError)
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.Response[any]{nil})
		return
	}

	response.Data = student
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetAllStudents(w http.ResponseWriter, _ *http.Request) {
	dbInstance := db.Connect()
	response := models.Response[[]models.Student]{[]models.Student{}}

	query := "SELECT student_id, student_name, course, present_date FROM student"

	rows, err := dbInstance.Query(query)
	if err != nil {
		apiresponse.SendErrorResponse(w, err, http.StatusInternalServerError, apiresponse.MessageInvalidQuery)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.StudentId, &student.StudentName, &student.Course, &student.PresentDate)
		if err != nil {
			apiresponse.SendErrorResponse(w, err, http.StatusInternalServerError, apiresponse.MessageScanError)
			return
		}

		response.Data = append(response.Data, student)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		apiresponse.SendErrorResponse(w, err, http.StatusInternalServerError, apiresponse.MessageResponseEncodingError)
		return
	}
}

package student

import (
	"docker-project-api/internal/db"
	"docker-project-api/internal/models"
	"docker-project-api/internal/util/apiresponse"
	"encoding/json"
	"errors"
	"github.com/lib/pq"
	"io"
	"net/http"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	dbInstance := db.Connect()
	student, err := parseBody(w, r)
	if err != nil {
		return
	}

	query := "INSERT INTO student (student_id, student_name, course) VALUES($1, $2, $3)"

	_, err = dbInstance.Exec(query, student.StudentId, student.StudentName, student.Course)
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		if pqErr.Code == "23505" {
			apiresponse.SendErrorResponse(w, pqErr, http.StatusConflict, "Student already exists")
			return
		}
	} else if err != nil {
		apiresponse.SendErrorResponse(w, err, http.StatusInternalServerError, apiresponse.MessageInvalidQuery)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.MessageResponse{"Student created successfully!"})

}

func parseBody(w http.ResponseWriter, r *http.Request) (models.StudentInsert, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		apiresponse.SendErrorResponse(w, err, http.StatusInternalServerError, apiresponse.MessageInternalError)
		return models.StudentInsert{}, errors.New("")
	}
	defer r.Body.Close()

	studentStruct := models.StudentInsert{}
	err = json.Unmarshal(body, &studentStruct)
	if err != nil || studentStruct.StudentName == "" || studentStruct.Course == "" || studentStruct.StudentId == "" {
		apiresponse.SendErrorResponse(w, nil, http.StatusBadRequest, `
			Provided body does not contain required fields. Make sure you are using application/json format 
			and 'studentName' and 'course' fields are present and non-empty`)
		return models.StudentInsert{}, errors.New("")
	}

	return studentStruct, nil
}

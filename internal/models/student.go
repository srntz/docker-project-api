package models

type Student struct {
	StudentId   string `json:"studentId"`
	StudentName string `json:"studentName"`
	Course      string `json:"course"`
	PresentDate string `json:"presentDate"`
}

type StudentInsert struct {
	StudentId   string `json:"studentId"`
	StudentName string `json:"studentName"`
	Course      string `json:"course"`
}

package models

type Student struct {
	StudentId   int64  `json:"studentId"`
	StudentName string `json:"studentName"`
	Course      string `json:"course"`
	PresentDate string `json:"presentDate"`
}

type StudentInsert struct {
	StudentName string `json:"studentName"`
	Course      string `json:"course"`
}

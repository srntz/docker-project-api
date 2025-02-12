package models

type Response[T any] struct {
	Data T `json:"data"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error struct {
		StatusCode int    `json:"status_code"`
		Message    string `json:"message"`
	} `json:"error"`
}

package models

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

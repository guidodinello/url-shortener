package models

import (
	"encoding/json"
	"errors"
)

var ErrNotFound = errors.New("Resource Not Found")
var ErrEmptyUrl = errors.New("Empty Url")
var ErrInvalidUrl = errors.New("Invalid Url")

type APIError struct {
	Message string `json:"error"`
}

func ApiJsonError(message string) string {
	errorResponse := APIError{Message: message}
	jsonResponse, _ := json.Marshal(errorResponse)
	return string(jsonResponse)
}

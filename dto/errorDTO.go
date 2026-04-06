package dto

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

var ErrBadRequest = errors.New("user bad request in json body")
var ErrFullNameIsEmpty = errors.New("full name of user is empty")
var ErrInvalidEmail = errors.New("invalid email")
var ErrPasswordIsEmpty = errors.New("password is empty")
var ErrRoleIsEmpty = errors.New("the role is empty")

type ErrorDTO struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func NewErrorDTO(message string) *ErrorDTO {
	return &ErrorDTO{
		Message: message,
		Time:    time.Now(),
	}
}

func ErrorBadRequest(err error, w http.ResponseWriter) {
	errDTO := NewErrorDTO(err.Error())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(errDTO); err != nil {
		log.Printf("(SERVER ERROR) failed to encode error response: %v", err)
	}
}

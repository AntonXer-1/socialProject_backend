package dto

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

var ErrBadRequest = errors.New("user bad request in json body")
var ErrFullNameIsEmpty = errors.New("full name of user is empty")
var ErrInvalidEmail = errors.New("invalid email")
var ErrPasswordIsEmpty = errors.New("password is empty")
var ErrRoleIsEmpty = errors.New("the role is empty")

type ErrorDTO struct {
	message string
	time    time.Time
}

func NewErrorDTO(message string) *ErrorDTO {
	return &ErrorDTO{
		message: message,
		time:    time.Now(),
	}
}

func (e *ErrorDTO) ErrorDTOToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func SendError(err error, w http.ResponseWriter, r *http.Request) {
	errDTO := NewErrorDTO(err.Error())
	if err == ErrBadRequest || err == ErrFullNameIsEmpty ||
		err == ErrInvalidEmail || err == ErrPasswordIsEmpty ||
		err == ErrRoleIsEmpty {
		http.Error(w, errDTO.ErrorDTOToString(), http.StatusBadRequest)
	} else {
		http.Error(w, errDTO.ErrorDTOToString(), http.StatusInternalServerError)
	}
}

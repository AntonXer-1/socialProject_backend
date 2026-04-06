package httpserver

import (
	"encoding/json"
	"log"
	"net/http"
	"socialProject_backend/dto"
	"socialProject_backend/user"
)

/*
pattern: /user
method: POST
info: JSON in HTTP body request

succeed:
  - status code: 201 Created
  - response body: JSON represented created user

failed:
  - status code: 400, 409, 500, ...
  - response body: JSON with error + time
*/
func HandleRegistration(w http.ResponseWriter, r *http.Request) {
	userDTO := dto.UserDTO{}
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		dto.ErrorBadRequest(err, w)
		return
	}
	if err := userDTO.ValidationForRegistration(); err != nil {
		dto.ErrorBadRequest(err, w)
		return
	}
	newUser := user.NewUser(userDTO.FullName,
		userDTO.Email, userDTO.Password,
		userDTO.Phone, userDTO.PhotoURL, userDTO.Role)

	/* логика добавления user в БД */
	_ = newUser // заглушка, пока нет реализации с БД

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	ToResponse := dto.UserResponseDTO{
		ID:       0, // id подтягивается из БД, пока что захардкодил
		FullName: userDTO.FullName,
		Email:    userDTO.Email,
	}
	if err := json.NewEncoder(w).Encode(ToResponse); err != nil {
		log.Printf("(SERVER ERROR) failed to encode user response: %v", err)
		return
	}
}

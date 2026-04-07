package httpserver

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"socialProject_backend/database"
	"socialProject_backend/dto"
	"socialProject_backend/user"

	"github.com/jackc/pgx/v5"
)

type HttpHandler struct {
	conn *pgx.Conn
	ctx  context.Context
}

func CreateHttpHandler(conn *pgx.Conn, ctx context.Context) *HttpHandler {
	return &HttpHandler{conn: conn,
		ctx: ctx,
	}
}

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
func (h *HttpHandler) HandleRegistration(w http.ResponseWriter, r *http.Request) {
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

	// check user in database
	status, err := database.CheckUser(h.conn, h.ctx, newUser.Email, newUser.Role)
	if err != nil {
		http.Error(w, "user verification error", http.StatusInternalServerError)
	}

	if status {
		response := dto.CreatAnser(false, "This user exists in database")
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(response)
		return
	}

	/* логика добавления user в БД */
	er := database.AddUser(h.conn, h.ctx, newUser)
	if er != nil {
		http.Error(w, "Error adding user", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	ToResponse := dto.UserResponseDTO{
		FullName: userDTO.FullName,
		Email:    userDTO.Email,
	}
	if err := json.NewEncoder(w).Encode(ToResponse); err != nil {
		log.Printf("(SERVER ERROR) failed to encode user response: %v", err)
		return
	}
}

/*
pattern: /authorization
method: GET
info: JSON in HTTP body request

succeed:
  - status code: 202 Created
  - response body: JSON User authenticated status

failed:
  - status code: 400, 409, 500, ...
  - response body: JSON with error + time
*/

func (h *HttpHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	ur := dto.UserDTO{}
	if err := json.NewDecoder(r.Body).Decode(&ur); err != nil {
		dto.ErrorBadRequest(err, w)
	}

	ans, err := database.CheckUser(h.conn, h.ctx, ur.Email, ur.Role)
	if err != nil {
		http.Error(w, "user verification error", http.StatusInternalServerError)
	}

	response := dto.CreatAnser(ans, "User authenticated status")

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

package httpserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	handler *HTTPHandler
}

func CreateServer(handler *HTTPHandler) *server {
	return &server{
		handler: handler,
	}
}

func (s *server) StartHTTPServer() {
	router := mux.NewRouter()

	router.Path("/user/registration").Methods("POST").HandlerFunc(s.handler.HandleRegistration)
	router.Path("/user/authorization").Methods("GET").HandlerFunc(s.handler.HandleAuthorization)

	http.ListenAndServe(":8080", router)
}

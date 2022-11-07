package server

import (
	"log"
	"net/http"
	"solution/service"

	"github.com/go-chi/chi/v5"
)

type server struct {
	router  *chi.Mux
	handler *handler
}

func NewServer(s service.Service) *server {
	return &server{router: chi.NewRouter(), handler: NewHandler(s)}
}

func (s *server) Start() error {
	log.Println("starting server...")

	s.handler.InitHandlers(s.router)

	err := http.ListenAndServe(":8080", s.router)
	if err != nil {
		return err
	}

	return nil
}

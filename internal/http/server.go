package http

import (
	"net/http"

	"github.com/badsector998/pos/internal/app"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	App *app.App
}

func NewServer(app *app.App) *Server {
	return &Server{
		App: app,
	}
}

func (s *Server) Run() error {
	r := chi.NewRouter()

	return http.ListenAndServe(":3333", r)
}

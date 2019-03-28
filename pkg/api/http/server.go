package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net"
	"net/http"
	"os"
	"time"
)

type Server struct {
	listener net.Listener
	Handler  http.Handler
	Address  string
}

func NewServer() (s *Server) {
	s = &Server{
		Address: os.Getenv("HTTP_PORT"),
	}
}

func routes() (r *chi.Mux) {
	r = chi.NewRouter()
	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Mount()
		})
	})
	return
}
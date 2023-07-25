package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func NewServer(port string) *http.Server {
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Mount("/", NewWebServer())
	r.Mount("/api", NewAPIServer())

	return &http.Server{
		Addr:    port,
		Handler: r,
	}

}

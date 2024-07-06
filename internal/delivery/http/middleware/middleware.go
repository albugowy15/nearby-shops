package middleware

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Setup(app *chi.Mux) {
	app.Use(middleware.Logger)
	app.Use(middleware.RequestID)
	app.Use(middleware.Recoverer)
	app.Use(middleware.Compress(5, "text/html", "application/json"))
}

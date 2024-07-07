package route

import (
	"github.com/albugowy15/nearby-shops/internal/delivery/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RouteConfig struct {
	App            *chi.Mux
	ShopController *http.ShopController
}

func (rc *RouteConfig) Setup() {
	rc.App.Use(middleware.Logger)
	rc.App.Use(middleware.RequestID)
	rc.App.Use(middleware.Recoverer)
	rc.App.Use(middleware.Compress(5, "text/html", "application/json"))

	rc.App.Route("/v1", func(r chi.Router) {
		r.Get("/shops", rc.ShopController.Search)
		r.Get("/shops/{shopId}", rc.ShopController.Get)
		r.Post("/shops", rc.ShopController.Create)
		r.Put("/shops/{shopId}", rc.ShopController.Update)
		r.Delete("/shops/{shopId}", rc.ShopController.Delete)
	})
}

package main

import (
	"fmt"
	"log"
	httpx "net/http"

	"github.com/albugowy15/nearby-shops/internal/config"
	"github.com/albugowy15/nearby-shops/internal/delivery/http"
	"github.com/albugowy15/nearby-shops/internal/delivery/http/route"
	"github.com/albugowy15/nearby-shops/internal/repository"
	"github.com/albugowy15/nearby-shops/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func main() {
	viperConfig := config.NewViper()
	validatorConfig := config.NewValidator(viperConfig)
	db := config.NewDatabase(viperConfig)

	// setup repositories
	shopRepository := repository.NewShopRepository(db)

	// setup use cases
	shopUseCase := usecase.NewShopUseCase(db, shopRepository)

	// setup controller
	shopController := http.NewShopController(shopUseCase, validatorConfig)

	app := chi.NewMux()
	route := route.RouteConfig{
		App:            app,
		ShopController: shopController,
	}
	route.Setup()

	address := fmt.Sprintf("localhost:%s", viperConfig.GetString("PORT"))
	log.Printf("Server running on port %s", viperConfig.GetString("PORT"))
	httpx.ListenAndServe(address, app)
}

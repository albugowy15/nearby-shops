package main

import (
	"fmt"
	"log"
	httpx "net/http"

	"github.com/albugowy15/nearby-shops/docs"
	"github.com/albugowy15/nearby-shops/internal/config"
	"github.com/albugowy15/nearby-shops/internal/delivery/http"
	"github.com/albugowy15/nearby-shops/internal/delivery/http/route"
	"github.com/albugowy15/nearby-shops/internal/repository"
	"github.com/albugowy15/nearby-shops/internal/usecase"
	"github.com/go-chi/chi/v5"
)

//	@title			Nearby Shops Swagger Documentation
//	@version		1.0
//	@description	This is swagger documentation for Nearby Shops REST API.

//	@contact.name	Mohamad Kholid Bughowi
//	@contact.url	https://bughowi.com
//	@contact.email	kholidbughowi@gmail.com

//	@BasePath	/v1

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
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

  port := viperConfig.GetString("PORT")
  baseUrl := viperConfig.GetString("BASE_URL")

  docs.SwaggerInfo.Host = baseUrl
	address := fmt.Sprintf(":%s", port)
	log.Printf("Server running on port %s", port)
	httpx.ListenAndServe(address, app)
}

package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/albugowy15/nearby-shops/internal/config"
	"github.com/albugowy15/nearby-shops/internal/model"
	"github.com/albugowy15/nearby-shops/internal/usecase"
	"github.com/albugowy15/nearby-shops/misc/httputils"
)

type ShopController struct {
	Validator *config.ValidatorConfig
	UseCase   *usecase.ShopUseCase
}

func NewShopController(useCase *usecase.ShopUseCase, validator *config.ValidatorConfig) *ShopController {
	return &ShopController{
		UseCase:   useCase,
		Validator: validator,
	}
}

func (c *ShopController) Create(w http.ResponseWriter, r *http.Request) {
	createShopRequest := &model.CreateShopRequest{}
	err := httputils.GetBody(r, createShopRequest)
	if err != nil {
		httputils.SendError(w, httputils.ErrDecodeJsonBody, http.StatusBadRequest)
		return
	}
	err = c.Validator.ValidateStruct(createShopRequest)
	if err != nil {
		httputils.SendError(w, err, http.StatusBadRequest)
		return
	}
	err = c.UseCase.Create(createShopRequest)
	if err != nil {
		ucErr := err.(usecase.UseCaseError)
		httputils.SendError(w, ucErr, ucErr.Code)
		return
	}
	httputils.SendMessage(w, "shop created", http.StatusCreated)
}

func (c *ShopController) Get(w http.ResponseWriter, r *http.Request) {
	shopIdPath := r.PathValue("shopId")
	if len(shopIdPath) == 0 {
		httputils.SendError(w, errors.New("missing shopId path value"), http.StatusBadRequest)
		return
	}
	shopId, err := strconv.ParseInt(shopIdPath, 10, 64)
  if err != nil {
		httputils.SendError(w, errors.New("shopId must be integer"), http.StatusBadRequest)
    return
  }
	getShopRequest := &model.GetShopRequest{
		ID: shopId,
	}
	getShopResponse, err := c.UseCase.Get(getShopRequest)
	if err != nil {
		ucErr := err.(usecase.UseCaseError)
		httputils.SendError(w, ucErr, ucErr.Code)
		return
	}
	httputils.SendData(w, getShopResponse, http.StatusCreated)
}

func (c *ShopController) Delete(w http.ResponseWriter, r *http.Request) {
	shopIdPath := r.PathValue("shopId")
	if len(shopIdPath) == 0 {
		httputils.SendError(w, errors.New("missing shopId path value"), http.StatusBadRequest)
		return
	}
	shopId, err := strconv.ParseInt(shopIdPath, 10, 64)
  if err != nil {
		httputils.SendError(w, errors.New("shopId must be integer"), http.StatusBadRequest)
    return
  }
	deleteShopRequest := &model.DeleteShopRequest{
		ID: shopId,
	}
	err = c.UseCase.Delete(deleteShopRequest)
	if err != nil {
		ucErr := err.(usecase.UseCaseError)
		httputils.SendError(w, ucErr, ucErr.Code)
		return
	}
	httputils.SendMessage(w, "shop deleted", http.StatusCreated)
}

func (c *ShopController) Update(w http.ResponseWriter, r *http.Request) {
	updateShopRequest := &model.UpdateShopRequest{}
	err := httputils.GetBody(r, updateShopRequest)
	if err != nil {
		httputils.SendError(w, httputils.ErrDecodeJsonBody, http.StatusBadRequest)
		return
	}
	err = c.Validator.ValidateStruct(updateShopRequest)
	if err != nil {
		httputils.SendError(w, err, http.StatusBadRequest)
		return
	}
	err = c.UseCase.Update(updateShopRequest)
	if err != nil {
		ucErr := err.(usecase.UseCaseError)
		httputils.SendError(w, ucErr, ucErr.Code)
		return
	}
	httputils.SendMessage(w, "shop updated", http.StatusCreated)
}

func (c *ShopController) Search(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	maxDistanceRaw := queryParams.Get("maxDistance")
	if len(maxDistanceRaw) == 0 {
		httputils.SendError(w, errors.New("maxDistance param required"), http.StatusBadRequest)
		return
	}
	longitudeRaw := queryParams.Get("lon")
	if len(longitudeRaw) == 0 {
		httputils.SendError(w, errors.New("lon param required"), http.StatusBadRequest)
		return
	}
	latitudeRaw := queryParams.Get("lat")
	if len(latitudeRaw) == 0 {
		httputils.SendError(w, errors.New("lat param required"), http.StatusBadRequest)
		return
	}

	maxDistance, err := strconv.ParseInt(maxDistanceRaw, 10, 64)
	if err != nil {
		httputils.SendError(w, errors.New("maxDistance param must be integer"), http.StatusBadRequest)
		return
	}
	longitude, err := strconv.ParseFloat(longitudeRaw, 64)
	if err != nil {
		httputils.SendError(w, errors.New("lon param must be integer"), http.StatusBadRequest)
		return
	}
	latitude, err := strconv.ParseFloat(latitudeRaw, 64)
	if err != nil {
		httputils.SendError(w, errors.New("lat param must be integer"), http.StatusBadRequest)
		return
	}
	data, err := c.UseCase.Search(&model.SearchShopRequest{
		MaxDistance: maxDistance,
		Longitude:   longitude,
		Latitude:    latitude,
	})
	if err != nil {
		ucErr := err.(usecase.UseCaseError)
		httputils.SendError(w, ucErr, ucErr.Code)
		return
	}
	httputils.SendData(w, data, http.StatusCreated)
}

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

// ShopController.Create godoc
//
//	@Summary		Create a new shop
//	@Description	Create a new shop
//	@Tags			Shop
//	@Accept			json
//	@Produce		json
//	@Param			body	body		model.CreateShopRequest	true	"Create shop request body"
//	@Success		201		{object}	model.MessageResponse
//	@Failure		400		{object}	model.ErrorResponse
//	@Failure		500		{object}	model.ErrorResponse
//	@Router			/shops [post]
func (c *ShopController) Create(w http.ResponseWriter, r *http.Request) {
	createShopRequest := &model.CreateShopRequest{}

	if err := httputils.GetBody(r, createShopRequest); err != nil {
		httputils.SendError(w, httputils.ErrDecodeJsonBody, http.StatusBadRequest)
		return
	}

	if err := c.Validator.ValidateStruct(createShopRequest); err != nil {
		httputils.SendError(w, err, http.StatusBadRequest)
		return
	}

	descLen := len(createShopRequest.Description)
	if descLen != 0 && (descLen < 30 || descLen > 300) {
		httputils.SendError(w, errors.New("description must be between 30 to 300 characters"), http.StatusBadRequest)
		return
	}

	if err := c.UseCase.Create(createShopRequest); err != nil {
		ucErr := err.(usecase.UseCaseError)
		httputils.SendError(w, ucErr, ucErr.Code)
		return
	}
	httputils.SendMessage(w, "shop created", http.StatusCreated)
}

// ShopController.Get godoc
//
//	@Summary		Get a shop details
//	@Description	Get detail information about a shop
//	@Tags			Shop
//	@Accept			json
//	@Produce		json
//	@Param			shopId	path		string	true	"Shop ID"
//	@Success		200		{object}	model.DataResponse{data=model.ShopLongResponse}
//	@Failure		400		{object}	model.ErrorResponse
//	@Failure		404		{object}	model.ErrorResponse
//	@Failure		500		{object}	model.ErrorResponse
//	@Router			/shops/{shopId} [get]
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

// ShopController.Delete godoc
//
//	@Summary		Delete a shop
//	@Description	Delete a shop
//	@Tags			Shop
//	@Accept			json
//	@Produce		json
//	@Param			shopId	path		string	true	"Shop ID"
//	@Success		200		{object}	model.MessageResponse
//	@Failure		400		{object}	model.ErrorResponse
//	@Failure		404		{object}	model.ErrorResponse
//	@Failure		500		{object}	model.ErrorResponse
//	@Router			/shops/{shopId} [delete]
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

	if err := c.UseCase.Delete(deleteShopRequest); err != nil {
		ucErr := err.(usecase.UseCaseError)
		httputils.SendError(w, ucErr, ucErr.Code)
		return
	}
	httputils.SendMessage(w, "shop deleted", http.StatusCreated)
}

// ShopController.Update godoc
//
//	@Summary		Update a shop
//	@Description	Update a shop
//	@Tags			Shop
//	@Accept			json
//	@Produce		json
//	@Param			shopId	path		string					true	"Shop ID"
//	@Param			body	body		model.UpdateShopRequest	true	"Update shop request body"
//	@Success		200		{object}	model.MessageResponse
//	@Failure		400		{object}	model.ErrorResponse
//	@Failure		404		{object}	model.ErrorResponse
//	@Failure		500		{object}	model.ErrorResponse
//	@Router			/shops/{shopId} [put]
func (c *ShopController) Update(w http.ResponseWriter, r *http.Request) {
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
	updateShopRequest := &model.UpdateShopRequest{}
	if err := httputils.GetBody(r, updateShopRequest); err != nil {
		httputils.SendError(w, httputils.ErrDecodeJsonBody, http.StatusBadRequest)
		return
	}
	if err := c.Validator.ValidateStruct(updateShopRequest); err != nil {
		httputils.SendError(w, err, http.StatusBadRequest)
		return
	}
	descLen := len(updateShopRequest.Description)
	if descLen != 0 && (descLen < 30 || descLen > 300) {
		httputils.SendError(w, errors.New("description must be between 30 to 300 characters"), http.StatusBadRequest)
		return
	}
	if err := c.UseCase.Update(updateShopRequest, shopId); err != nil {
		ucErr := err.(usecase.UseCaseError)
		httputils.SendError(w, ucErr, ucErr.Code)
		return
	}
	httputils.SendMessage(w, "shop updated", http.StatusCreated)
}

// ShopController.Search godoc
//
//	@Summary		Search nearby shops
//	@Description	Search nearby shops
//	@Tags			Shop
//	@Accept			json
//	@Produce		json
//	@Param			maxDistance	query		integer	true	"Search nearby shops by maxDistance"
//	@Param			lon			query		number	true	"Search nearby shops by Longitude"
//	@Param			lat			query		number	true	"Search nearby shops by Latitude"
//	@Success		200			{object}	model.DataResponse{data=[]model.ShopResponse}
//	@Failure		400			{object}	model.ErrorResponse
//	@Failure		500			{object}	model.ErrorResponse
//	@Router			/shops [get]
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

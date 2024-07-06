package http

import (
	"errors"
	"net/http"

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
	httputils.SendData(w, "shop created", http.StatusCreated)
}

func (c *ShopController) Get(w http.ResponseWriter, r *http.Request) {
	shopId := r.PathValue("shopId")
	if len(shopId) == 0 {
		httputils.SendError(w, errors.New("missing shopId path value"), http.StatusBadRequest)
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
	shopId := r.PathValue("shopId")
	if len(shopId) == 0 {
		httputils.SendError(w, errors.New("missing shopId path value"), http.StatusBadRequest)
		return
	}
	deleteShopRequest := &model.DeleteShopRequest{
		ID: shopId,
	}
	err := c.UseCase.Delete(deleteShopRequest)
	if err != nil {
		ucErr := err.(usecase.UseCaseError)
		httputils.SendError(w, ucErr, ucErr.Code)
		return
	}
	httputils.SendData(w, "shop deleted", http.StatusCreated)
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

func (c *ShopController) List(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hello"))
}

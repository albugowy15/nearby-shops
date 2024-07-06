package usecase

import (
	"net/http"

	"github.com/albugowy15/nearby-shops/internal/model"
	"github.com/albugowy15/nearby-shops/internal/repository"
	"github.com/jmoiron/sqlx"
)

type ShopUseCase struct {
	DB             *sqlx.DB
	ShopRepository *repository.ShopRepository
}

func NewShopUseCase(db *sqlx.DB, shopRepository *repository.ShopRepository) *ShopUseCase {
	return &ShopUseCase{
		DB:             db,
		ShopRepository: shopRepository,
	}
}

func (uc *ShopUseCase) Create(request *model.CreateShopRequest) error {
	return NewUseCaseError("hello error", http.StatusGatewayTimeout)
}

func (uc *ShopUseCase) Get(request *model.GetShopRequest) (*model.ShopLongResponse, error) {
	return &model.ShopLongResponse{}, nil
}

func (uc *ShopUseCase) Search(request *model.ListShopRequest) ([]model.ShopResponse, error) {
	shops := []model.ShopResponse{}
	return shops, nil
}

func (uc *ShopUseCase) Update(request *model.UpdateShopRequest) error {
	return nil
}

func (uc *ShopUseCase) Delete(request *model.DeleteShopRequest) error {
	return nil
}

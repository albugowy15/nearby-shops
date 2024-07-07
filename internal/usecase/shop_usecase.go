package usecase

import (
	"database/sql"
	"net/http"

	"github.com/albugowy15/nearby-shops/internal/entity"
	"github.com/albugowy15/nearby-shops/internal/model"
	"github.com/albugowy15/nearby-shops/internal/repository"
	"github.com/albugowy15/nearby-shops/misc/postgis"
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
	shopValue := entity.Shop{
		Name: request.Name,
		City: request.City,
	}
	if len(request.Description) != 0 {
		shopValue.Description = sql.NullString{
			String: request.Description,
			Valid:  true,
		}
	}
	locationPoint := postgis.NewFromPoint(postgis.Point{
		Lon: request.Longitude,
		Lat: request.Latitude,
	})

	shopValue.Location = locationPoint.Ewkt()
	err := uc.ShopRepository.Insert(shopValue)
	if err != nil {
		return NewUseCaseError("error inserting data into database", http.StatusInternalServerError)
	}
	return nil
}

func (uc *ShopUseCase) Get(request *model.GetShopRequest) (*model.ShopLongResponse, error) {
	return &model.ShopLongResponse{}, nil
}

func (uc *ShopUseCase) Search(request *model.SearchShopRequest) ([]model.ShopResponse, error) {
	shops := []model.ShopResponse{}
	point := postgis.NewFromPoint(postgis.Point{
		Lon: request.Longitude,
		Lat: request.Latitude,
	})
	shopRows, err := uc.ShopRepository.FilterByLocation(request.MaxDistance, *point)
	if err != nil {
		return shops, NewUseCaseError("error fething data from database", http.StatusInternalServerError)
	}

	for _, row := range shopRows {
		item := model.ShopResponse{
			ID:       row.ID,
			Name:     row.Name,
			Distance: row.Distance,
		}
		point, err := postgis.NewFromEwkt(row.Location)
		if err != nil {
			return shops, NewUseCaseError("error read geograpy data type database", http.StatusInternalServerError)
		}
		item.Longitude = point.Point.Lon
		item.Latitude = point.Point.Lat
		shops = append(shops, item)
	}
	return shops, nil
}

func (uc *ShopUseCase) Update(request *model.UpdateShopRequest) error {
	return nil
}

func (uc *ShopUseCase) Delete(request *model.DeleteShopRequest) error {
	return nil
}

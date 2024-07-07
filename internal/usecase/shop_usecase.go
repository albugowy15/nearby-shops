package usecase

import (
	"database/sql"
	"fmt"
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

	if err := uc.ShopRepository.Insert(shopValue); err != nil {
		return NewUseCaseError("error inserting data into database", http.StatusInternalServerError)
	}
	return nil
}

func (uc *ShopUseCase) Get(request *model.GetShopRequest) (*model.ShopLongResponse, error) {
	findResult, err := uc.ShopRepository.FindById(request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewUseCaseError(fmt.Sprintf("shop with id %d not found", request.ID), http.StatusNotFound)
		}
		return nil, NewUseCaseError("error fetching data from database", http.StatusInternalServerError)
	}

	point, err := postgis.NewFromEwkt(findResult.Location)
	if err != nil {
		return nil, NewUseCaseError("error read location data from database", http.StatusInternalServerError)
	}
	shop := &model.ShopLongResponse{
		ID:        findResult.ID,
		Name:      findResult.Name,
		Longitude: point.Point.Lon,
		Latitude:  point.Point.Lat,
		City:      findResult.City,
		CreatedAt: findResult.CreatedAt,
	}
	if findResult.Description.Valid {
		shop.Description = findResult.Description.String
	}

	return shop, nil
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

func (uc *ShopUseCase) Update(request *model.UpdateShopRequest, id int64) error {
	isRowExist, err := uc.ShopRepository.CheckRowExist(id)
	if err != nil {
		return NewUseCaseError("error fetching data from database", http.StatusInternalServerError)
	}
	if !isRowExist {
		return NewUseCaseError(fmt.Sprintf("shop with id %d not found", id), http.StatusNotFound)
	}
	point := postgis.NewFromPoint(postgis.Point{Lon: request.Longitude, Lat: request.Latitude})
	value := entity.Shop{
		Name:     request.Name,
		City:     request.City,
		Location: point.Ewkt(),
	}
	if len(request.Description) != 0 {
		value.Description = sql.NullString{Valid: true, String: request.Description}
	}

	if err := uc.ShopRepository.Update(id, value); err != nil {
		return NewUseCaseError("error update row from database", http.StatusInternalServerError)
	}
	return nil
}

func (uc *ShopUseCase) Delete(request *model.DeleteShopRequest) error {
	isRowExist, err := uc.ShopRepository.CheckRowExist(request.ID)
	if err != nil {
		return NewUseCaseError("error fetching data from database", http.StatusInternalServerError)
	}
	if !isRowExist {
		return NewUseCaseError(fmt.Sprintf("shop with id %d not found", request.ID), http.StatusNotFound)
	}
	if err := uc.ShopRepository.Delete(request.ID); err != nil {
		return NewUseCaseError("error delete row from database", http.StatusInternalServerError)
	}
	return nil
}

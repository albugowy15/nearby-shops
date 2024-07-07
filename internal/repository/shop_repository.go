package repository

import (
	"log"
	"time"

	"github.com/albugowy15/nearby-shops/internal/entity"
	"github.com/albugowy15/nearby-shops/misc/postgis"
	"github.com/jmoiron/sqlx"
)

type ShopRepository struct {
	DB *sqlx.DB
}

func NewShopRepository(db *sqlx.DB) *ShopRepository {
	return &ShopRepository{
		DB: db,
	}
}

type FilterByLocationResult struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	City      string    `db:"city"`
	Location  string    `db:"location"`
	Distance  float64   `db:"distance"`
	CreatedAt time.Time `db:"created_at"`
}

func (r *ShopRepository) FilterByLocation(maxDistance int64, point postgis.PostGisGeo) ([]FilterByLocationResult, error) {
	shops := []FilterByLocationResult{}
	err := r.DB.Select(
		&shops,
		`
    select id, name, city, st_asewkt(location) as location, created_at,
    ST_Distance($1::geography, location::geography) as distance
    from shops
    where st_dwithin(location, $2::geography, $3)
    `,
		point.Ewkt(), point.Ewkt(), maxDistance,
	)
	if err != nil {
		log.Printf("db err: %v", err)
		return shops, err
	}
	return shops, nil
}

func (r *ShopRepository) Insert(value entity.Shop) error {
	_, err := r.DB.Exec(
		"insert into shops(name, city, description, location) values ($1, $2, $3, $4)",
		value.Name, value.City, value.Description, value.Location,
	)
	if err != nil {
		log.Printf("db err: %v", err)
		return err
	}
	return nil
}

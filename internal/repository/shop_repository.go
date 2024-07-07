package repository

import (
	"database/sql"
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
	CreatedAt time.Time `db:"created_at"`
	Name      string    `db:"name"`
	City      string    `db:"city"`
	Location  string    `db:"location"`
	ID        int64     `db:"id"`
	Distance  float64   `db:"distance"`
}

type FindByIdResult struct {
	CreatedAt   time.Time      `db:"created_at"`
	Name        string         `db:"name"`
	City        string         `db:"city"`
	Location    string         `db:"location"`
	Description sql.NullString `db:"description"`
	ID          int64          `db:"id"`
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
		`
    insert into shops(name, city, description, location) values ($1, $2, $3, $4)
    `,
		value.Name, value.City, value.Description, value.Location,
	)
	if err != nil {
		log.Printf("db err: %v", err)
		return err
	}
	return nil
}

func (r *ShopRepository) FindById(shopId int64) (FindByIdResult, error) {
	result := FindByIdResult{}
	err := r.DB.Get(
		&result,
		`
    select id, name, city, st_asewkt(location) as location, description, created_at
    from shops
    where id = $1
    `,
		shopId,
	)
	if err != nil {
		log.Printf("db err: %v", err)
	}
	return result, err
}

func (r *ShopRepository) CheckRowExist(shopId int64) (bool, error) {
	var resultId int64
	err := r.DB.Get(&resultId, "select id from shops where id = $1", shopId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *ShopRepository) Delete(shopId int64) error {
	_, err := r.DB.Exec(
		`delete from shops where id = $1`,
		shopId,
	)
	if err != nil {
		log.Printf("db err: %v", err)
	}
	return err
}

func (r *ShopRepository) Update(shopId int64, value entity.Shop) error {
	_, err := r.DB.Exec(
		`
    update shops
    set name = $1, city = $2, location = $3, description = $4 
    where id = $5
    `,
		value.Name, value.City, value.Location, value.Description, shopId,
	)
	if err != nil {
		log.Printf("db err: %v", err)
	}
	return err
}

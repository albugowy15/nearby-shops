package repository

import "github.com/jmoiron/sqlx"

type ShopRepository struct {
	DB *sqlx.DB
}

func NewShopRepository(db *sqlx.DB) *ShopRepository {
	return &ShopRepository{
		DB: db,
	}
}

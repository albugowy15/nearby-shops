package model

import "time"

type ShopResponse struct {
	Name      string  `json:"name"`
	ID        int64   `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Distance  float64 `json:"distance"`
}

type ShopLongResponse struct {
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	City        string    `json:"city"`
	ID          int64     `json:"id"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
}

type CreateShopRequest struct {
	Name        string  `json:"name" validate:"required,max=100"`
	Description string  `json:"description,omitempty"`
	City        string  `json:"city" validate:"required,max=50"`
	Latitude    float64 `json:"latitude" validate:"required"`
	Longitude   float64 `json:"longitude" validate:"required"`
}

type UpdateShopRequest struct {
	Name        string  `json:"name" validate:"required,max=100"`
	Description string  `json:"description,omitempty"`
	City        string  `json:"city" validate:"required,max=50"`
	Latitude    float64 `json:"latitude" validate:"required"`
	Longitude   float64 `json:"longitude" validate:"required"`
}

type DeleteShopRequest struct {
	ID int64 `json:"id"`
}

type GetShopRequest struct {
	ID int64 `json:"id"`
}

type SearchShopRequest struct {
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	MaxDistance int64   `json:"max_distance"`
}

package model

type ShopResponse struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Distance  int64   `json:"distance"`
}

type ShopLongResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Distance    int64   `json:"distance"`
	CreatedAt   int64   `json:"created_at"`
}

type CreateShopRequest struct {
	Name        string  `json:"name" validate:"required,max=100"`
	Description string  `json:"description,omitempty"`
	Latitude    float64 `json:"latitude" validate:"required"`
	Longitude   float64 `json:"longitude" validate:"required"`
}

type UpdateShopRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type DeleteShopRequest struct {
	ID string `json:"id"`
}

type GetShopRequest struct {
	ID string `json:"id"`
}

type ListShopRequest struct {
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	MaxDistance int64   `json:"max_distance"`
}

package model

type ErrorResponse struct {
	Error string `json:"error"`
}

type DataResponse struct {
	Data interface{} `json:"data"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

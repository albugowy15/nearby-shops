package httputils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/albugowy15/nearby-shops/internal/model"
)

var (
	ErrInternalServer = errors.New("internal server errror")
	ErrDecodeJsonBody = errors.New("error decode json body")
)

func GetBody(r *http.Request, v any) error {
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		log.Printf("error decode json body: %v", err)
		return ErrDecodeJsonBody
	}
	return nil
}

func Send(w http.ResponseWriter, res any, status int) {
	json, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
}

func SendError(w http.ResponseWriter, err error, status int) {
	res := model.ErrorResponse{
		Error: err.Error(),
	}
	Send(w, res, status)
}

func SendMessage(w http.ResponseWriter, message string, status int) {
	res := model.MessageResponse{
		Message: message,
	}
	Send(w, res, status)
}

func SendData(w http.ResponseWriter, data interface{}, status int) {
	res := model.DataResponse{
		Data: data,
	}
	Send(w, res, status)
}

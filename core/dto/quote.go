package dto

import (
	"encoding/json"
	"io"
	"log"

	"github.com/go-playground/validator/v10"
)

type QuoteRequest struct {
	Shipper        Shipper      `json:"shipper" validate:"required"`
	Recipient      Recipient    `json:"recipient" validate:"required"`
	Dispatchers    []Dispatcher `json:"dispatchers" validate:"required,dive,required"`
	SimulationType []int64      `json:"simulation_type" validate:"required"`
}

type Shipper struct {
	RegisteredNumber string `json:"registered_number" validate:"required`
	Token            string `json:"token" validate:"required`
	PlatformCode     string `json:"platform_code" validate:"required`
}

type Recipient struct {
	Type             int    `json:"type" validate:"required`
	RegisteredNumber string `json:"registered_number" validate:"required`
	StateInscription string `json:"state_inscription" validate:"required`
	Country          string `json:"country" validate:"required`
	Zipcode          int64  `json:"zipcode" validate:"required`
}

type Dispatcher struct {
	RegisteredNumber string   `json:"registered_number" validate:"required`
	Zipcode          int64    `json:"zipcode" validate:"required`
	TotalPrice       float64  `json:"total_price" validate:"required`
	Volumes          []Volume `json:"volumes" validate:"required,dive,required`
}

type Volume struct {
	Category      string  `json:"category" validate:"required`
	Amount        int64   `json:"amount" validate:"required`
	UnitaryWeight float64 `json:"unitary_weight" validate:"required`
	Sku           string  `json:"sku" validate:"required`
	Height        float64 `json:"height" validate:"required`
	Width         float64 `json:"width" validate:"required`
	Length        float64 `json:"length" validate:"required`
	UnitaryPrice  float64 `json:"unitary_price" validate:"required`
}

type DispatchersResponse struct {
	ID                         string `json:"id"`
	RequestID                  string `json:"request_id"`
	RegisteredNumberShipper    string `json:"registered_number_shipper"`
	RegisteredNumberDispatcher string `json:"registered_number_dispatcher"`
	ZipcodeOrigin              string `json:"zipcode_origin"`
}

type QuoteStore struct {
	ID       string  `db:"id"`
	Name     string  `db:"name"`
	Service  string  `db:"service"`
	Deadline string  `db:"deadline"`
	Price    float64 `db:"price"`
}

func FromJSONCreateQuoteRequest(body io.Reader) (*QuoteRequest, error) {
	quoteRequest := QuoteRequest{}
	if err := json.NewDecoder(body).Decode(&quoteRequest); err != nil {
		log.Println("Error decoding quote request")
		return nil, err
	}
	validate := validator.New()
	if err := validate.Struct(quoteRequest); err != nil {
		return nil, err
	}
	return &quoteRequest, nil
}

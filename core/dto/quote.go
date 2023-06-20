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
	RegisteredNumber string `json:"registered_number"`
	Token            string `json:"token"`
	PlatformCode     string `json:"platform_code"`
}

type Recipient struct {
	Type             int    `json:"type"`
	RegisteredNumber string `json:"registered_number"`
	StateInscription string `json:"state_inscription"`
	Country          string `json:"country"`
	Zipcode          int64  `json:"zipcode"`
}

type Dispatcher struct {
	RegisteredNumber string   `json:"registered_number"`
	Zipcode          int64    `json:"zipcode"`
	TotalPrice       float64  `json:"total_price"`
	Volumes          []Volume `json:"volumes"`
}

type Volume struct {
	Category      string  `json:"category"`
	Amount        int64   `json:"amount"`
	UnitaryWeight float64 `json:"unitary_weight"`
	Sku           string  `json:"sku"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
	UnitaryPrice  float64 `json:"unitary_price"`
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

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

type QuoteStore struct {
	ID       string  `db:"id"`
	Name     string  `db:"name"`
	Service  string  `db:"service"`
	Deadline int     `db:"deadline"`
	Price    float64 `db:"price"`
}

type DispatchersResponse struct {
	Dispatchers []DispatcherResponse `json:"dispatchers"`
}

type DispatcherResponse struct {
	ID                         string          `json:"id"`
	RequestID                  string          `json:"request_id"`
	RegisteredNumberShipper    string          `json:"registered_number_shipper"`
	RegisteredNumberDispatcher string          `json:"registered_number_dispatcher"`
	ZipcodeOrigin              int64           `json:"zipcode_origin"`
	Offers                     []OfferResponse `json:"offers"`
}

type OfferResponse struct {
	Offer                int                  `json:"offer"`
	TableReference       string               `json:"table_reference"`
	SimulationType       int64                `json:"simulation_type"`
	Carrier              CarrierResponse      `json:"carrier"`
	Service              string               `json:"service"`
	ServiceCode          string               `json:"service_code"`
	ServiceDescription   string               `json:"service_description"`
	DeliveryTime         DeliveryTimeResponse `json:"delivery_time"`
	Expiration           string               `json:"expiration"`
	CostPrice            float64              `json:"cost_price"`
	FinalPrice           float64              `json:"final_price"`
	Weights              WeightsResponse      `json:"weights"`
	OriginalDeliveryTime DeliveryTimeResponse `json:"original_delivery_time"`
}

type CarrierResponse struct {
	Name             string `json:"name"`
	RegisteredNumber string `json:"registered_number"`
	StateInscription string `json:"state_inscription"`
	Logo             string `json:"logo"`
	Reference        int    `json:"reference"`
	CompanyName      string `json:"company_name"`
}

type DeliveryTimeResponse struct {
	Days          int    `json:"days"`
	Hours         int    `json:"hours,omitempty"`
	Minutes       int    `json:"minutes,omitempty"`
	EstimatedDate string `json:"estimated_date"`
}

type WeightsResponse struct {
	Real  float64 `json:"real"`
	Cubed float64 `json:"cubed,omitempty"`
	Used  float64 `json:"used"`
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

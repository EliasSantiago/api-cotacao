package dto

import (
	"encoding/json"
	"io"
)

type CreateQuoteRequest struct {
	Name     string  `json:"name" validate:"required,name"`
	Service  string  `json:"service" validate:"required,service"`
	Deadline string  `json:"deadline" validate:"required,deadline"`
	Price    float64 `json:"price" validate:"required,price"`
}

type CreateQuoteResponse struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline string  `json:"deadline"`
	Price    float64 `json:"price"`
}

type CreateQuoteStore struct {
	ID       string  `db:"id"`
	Name     string  `db:"name"`
	Service  string  `db:"service"`
	Deadline string  `db:"deadline"`
	Price    float64 `db:"price"`
}

func FromJSONCreateQuoteRequest(body io.Reader) (*CreateQuoteRequest, error) {
	createQuoteRequest := CreateQuoteRequest{}
	if err := json.NewDecoder(body).Decode(&createQuoteRequest); err != nil {
		return nil, err
	}
	return &createQuoteRequest, nil
}

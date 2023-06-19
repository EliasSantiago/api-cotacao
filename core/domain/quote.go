package domain

import (
	"net/http"

	"github.com/EliasSantiago/api-cotacao/core/dto"
)

type Quote struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline string  `json:"deadline"`
	Price    float64 `json:"price"`
}

type QuoteService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Metrics(response http.ResponseWriter, request *http.Request)
}

type QuoteUseCase interface {
	Create(quoteRequest *dto.CreateQuoteRequest) error
	Metrics(metricsRequest *dto.MetricsRequestParms) (*MetricsResponse, error)
}

type QuoteRepository interface {
	Create(quoteRequest *dto.CreateQuoteStore) error
	Metrics(metricsRequest *dto.MetricsRequestParms) (*MetricsResponse, error)
}

package domain

import (
	"net/http"

	"github.com/EliasSantiago/api-cotacao/core/dto"
)

type Quote struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline int     `json:"deadline"`
	Price    float64 `json:"price"`
}

type QuoteResponse struct {
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline int     `json:"deadline"`
	Price    float64 `json:"price"`
}

type OffersResponse struct {
	Offers []QuoteResponse `json:"offers"`
}

type QuoteService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Metrics(response http.ResponseWriter, request *http.Request)
}

type QuoteUseCase interface {
	Create(quoteRequest *dto.QuoteRequest) (*OffersResponse, error)
	Metrics(metricsRequest *dto.MetricsRequestParms) (*MetricsResponse, error)
}

type QuoteRepository interface {
	Create(quoteRequest *dto.QuoteStore) (*Quote, error)
	Metrics(metricsRequest *dto.MetricsRequestParms) (*MetricsResponse, error)
}

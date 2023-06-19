package quoteservice

import (
	"encoding/json"
	"net/http"

	"github.com/EliasSantiago/api-cotacao/core/dto"
)

// @Summary Metrics
// @Description Metrics quote
// @Tags metrics
// @Accept  json
// @Produce  json
// @Param quote body dto.CreateQuoteResponse true "quote"
// @Success 200 {object} domain.Quote
// @Router /metrics [get]
func (service service) Metrics(response http.ResponseWriter, request *http.Request) {
	metricsRequest, err := dto.FromValueMetricsRequestParams(request)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	quotes, err := service.usecase.Metrics(metricsRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(quotes)
}

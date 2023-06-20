package quoteservice

import (
	"encoding/json"
	"net/http"

	"github.com/EliasSantiago/api-cotacao/adapter/http/errors"
	"github.com/EliasSantiago/api-cotacao/core/dto"
)

// @Summary Create new quote
// @Description Create new quote
// @Tags quote
// @Accept  json
// @Produce  json
// @Param quote body dto.CreateQuoteRequest true "quote"
// @Success 200 {object} domain.Quote
// @Router /quote [post]8
func (service service) Create(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	quoteRequest, err := dto.FromJSONCreateQuoteRequest(request.Body)
	if err != nil {
		errors.InvalidParameters()
		return
	}

	quote, err := service.usecase.Create(quoteRequest)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(quote)
}

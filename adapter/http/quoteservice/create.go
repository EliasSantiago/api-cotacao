package quoteservice

import (
	"encoding/json"
	"net/http"

	"github.com/EliasSantiago/api-cotacao/core/dto"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

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
		errResponse := ErrorResponse{
			Message: err.Error(),
		}
		response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(response).Encode(errResponse)
		return
	}

	quote, err := service.usecase.Create(quoteRequest)
	if err != nil {
		errResponse := ErrorResponse{
			Message: err.Error(),
		}
		response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(response).Encode(errResponse)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(quote)
}

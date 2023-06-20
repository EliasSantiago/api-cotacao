package quoteusecase

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/EliasSantiago/api-cotacao/core/domain"
	"github.com/EliasSantiago/api-cotacao/core/dto"
	"github.com/google/uuid"
)

func (usecase usecase) Create(quoteRequest *dto.QuoteRequest) (*domain.Quote, error) {
	payload, err := json.Marshal(quoteRequest)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("https://sp.freterapido.com/api/v3/quote/simulate", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var quoteResponse dto.DispatchersResponse
	if err := json.NewDecoder(resp.Body).Decode(&quoteResponse); err != nil {
		log.Println("Error decoding quote request")
		return nil, err
	}

	uuid := uuid.New()
	rand.Seed(time.Now().UnixNano())
	store := &dto.QuoteStore{
		ID:       strings.Replace(uuid.String(), "-", "", -1),
		Name:     "EXPRESSO FR",
		Service:  "Rodoviário",
		Deadline: "3",
		Price:    rand.Float64(),
	}

	quote, err := usecase.repository.Create(store)
	if err != nil {
		return nil, err
	}
	return quote, nil
}

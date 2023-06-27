package quoteusecase

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/EliasSantiago/api-cotacao/core/domain"
	"github.com/EliasSantiago/api-cotacao/core/dto"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func (usecase usecase) Create(quoteRequest *dto.QuoteRequest) (*domain.OffersResponse, error) {

	quoteRequest.Shipper.RegisteredNumber = viper.GetString("shipper.registered_number")
	quoteRequest.Shipper.Token = viper.GetString("shipper.token")
	quoteRequest.Shipper.PlatformCode = viper.GetString("shipper.platform_code")

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

	var offers []domain.QuoteResponse
	for _, dispatcher := range quoteResponse.Dispatchers {
		for _, offer := range dispatcher.Offers {
			uuid := uuid.New()
			store := &dto.QuoteStore{
				ID:       strings.Replace(uuid.String(), "-", "", -1),
				Name:     offer.Carrier.Name,
				Service:  offer.Service,
				Deadline: offer.DeliveryTime.Days,
				Price:    offer.FinalPrice,
			}
			_, err := usecase.repository.Create(store)
			if err != nil {
				return nil, err
			}

			quoteResponse := domain.QuoteResponse{
				Name:     offer.Carrier.Name,
				Service:  offer.Service,
				Deadline: offer.DeliveryTime.Days,
				Price:    offer.FinalPrice,
			}

			offers = append(offers, quoteResponse)
		}
	}
	response := &domain.OffersResponse{
		Offers: offers,
	}

	return response, nil
}

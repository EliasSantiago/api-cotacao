package quoteusecase

import (
	"strings"

	"github.com/EliasSantiago/api-cotacao/core/dto"
	"github.com/google/uuid"
)

func (usecase usecase) Create(quoteRequest *dto.CreateQuoteRequest) error {
	uuid := uuid.New()
	store := &dto.CreateQuoteStore{
		ID:       strings.Replace(uuid.String(), "-", "", -1),
		Name:     quoteRequest.Name,
		Service:  quoteRequest.Service,
		Deadline: quoteRequest.Deadline,
		Price:    quoteRequest.Price,
	}
	err := usecase.repository.Create(store)
	if err != nil {
		return err
	}
	return nil
}

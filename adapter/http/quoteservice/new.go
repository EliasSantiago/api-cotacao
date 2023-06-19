package quoteservice

import (
	"github.com/EliasSantiago/api-cotacao/core/domain"
)

type service struct {
	usecase domain.QuoteUseCase
}

func New(usecase domain.QuoteUseCase) domain.QuoteService {
	return &service{
		usecase: usecase,
	}
}

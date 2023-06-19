package quoteusecase

import "github.com/EliasSantiago/api-cotacao/core/domain"

type usecase struct {
	repository domain.QuoteRepository
}

func New(repository domain.QuoteRepository) domain.QuoteUseCase {
	return &usecase{
		repository: repository,
	}
}

package di

import (
	"github.com/EliasSantiago/api-cotacao/adapter/http/quoteservice"
	"github.com/EliasSantiago/api-cotacao/adapter/postgres"
	"github.com/EliasSantiago/api-cotacao/adapter/postgres/quoterepository"
	"github.com/EliasSantiago/api-cotacao/core/domain"
	"github.com/EliasSantiago/api-cotacao/core/domain/usecase/quoteusecase"
)

func ConfigQuoteDI(conn postgres.PoolInterface) domain.QuoteService {
	quoteRepository := quoterepository.New(conn)
	quoteUseCase := quoteusecase.New(quoteRepository)
	quoteService := quoteservice.New(quoteUseCase)
	return quoteService
}

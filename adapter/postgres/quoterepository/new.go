package quoterepository

import (
	"github.com/EliasSantiago/api-cotacao/adapter/postgres"
	"github.com/EliasSantiago/api-cotacao/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

func New(db postgres.PoolInterface) domain.QuoteRepository {
	return &repository{
		db: db,
	}
}

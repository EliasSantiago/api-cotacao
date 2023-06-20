package quoterepository

import (
	"context"

	"github.com/EliasSantiago/api-cotacao/core/domain"
	"github.com/EliasSantiago/api-cotacao/core/dto"
)

func (repository repository) Create(quoteRequest *dto.QuoteStore) (*domain.Quote, error) {
	ctx := context.Background()
	quote := domain.Quote{}
	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO quotes (id, name, service, deadline, price) VALUES ($1, $2, $3, $4, $5) returning *",
		quoteRequest.ID,
		quoteRequest.Name,
		quoteRequest.Service,
		quoteRequest.Deadline,
		quoteRequest.Price,
	).Scan(
		&quote.ID,
		&quote.Name,
		&quote.Service,
		&quote.Deadline,
		&quote.Price,
	)
	if err != nil {
		return nil, err
	}
	return &quote, nil
}

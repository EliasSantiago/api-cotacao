package quoterepository

import (
	"context"
	"fmt"

	"github.com/EliasSantiago/api-cotacao/core/domain"
	"github.com/EliasSantiago/api-cotacao/core/dto"
)

func (repository repository) Metrics(params *dto.MetricsRequestParms) (*domain.MetricsResponse, error) {
	ctx := context.Background()
	quotes := []domain.Quote{}

	rows, err := repository.db.Query(ctx, `SELECT * FROM quotes`)

	if err != nil {
		fmt.Println("Erro ao executar a consulta:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var q domain.Quote
		err := rows.Scan(&q.ID, &q.Name, &q.Service, &q.Deadline, &q.Price)
		if err != nil {
			fmt.Println("Erro ao mapear os resultados:", err)
			return nil, err
		}
		quotes = append(quotes, q)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Erro ao percorrer as linhas:", err)
		return nil, err
	}

	countByNames := make(map[string]int64)
	totalPriceByNames := make(map[string]float64)
	minPrice := quotes[0].Price
	maxPrice := quotes[0].Price

	for _, q := range quotes {
		countByNames[q.Name]++
		totalPriceByNames[q.Name] += q.Price

		if q.Price < minPrice {
			minPrice = q.Price
		}

		if q.Price > maxPrice {
			maxPrice = q.Price
		}
	}

	metrics := []domain.Metric{}

	for name, total := range totalPriceByNames {
		average := total / float64(countByNames[name])
		metric := domain.Metric{
			Name:       name,
			QtResults:  countByNames[name],
			TotalPrice: total,
			AvgPrice:   average,
		}
		metrics = append(metrics, metric)
	}

	return &domain.MetricsResponse{
		Metrics:  metrics,
		MinPrice: minPrice,
		MaxPrice: maxPrice,
	}, nil
}

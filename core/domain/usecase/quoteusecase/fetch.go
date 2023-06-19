package quoteusecase

import (
	"github.com/EliasSantiago/api-cotacao/core/domain"
	"github.com/EliasSantiago/api-cotacao/core/dto"
)

func (usecase usecase) Metrics(metricsRequest *dto.MetricsRequestParms) (*domain.MetricsResponse, error) {
	quotes, err := usecase.repository.Metrics(metricsRequest)

	if err != nil {
		return nil, err
	}

	return quotes, nil
}

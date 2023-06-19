package dto

import (
	"net/http"
	"strconv"
)

type MetricsRequestParms struct {
	LastQuotes int `json:"lastQuotes"`
}

func FromValueMetricsRequestParams(request *http.Request) (*MetricsRequestParms, error) {
	lastQuotes, _ := strconv.Atoi(request.FormValue("last_quotes"))
	metricsRequestParms := MetricsRequestParms{
		LastQuotes: lastQuotes,
	}

	return &metricsRequestParms, nil
}

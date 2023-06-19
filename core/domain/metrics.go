package domain

type Metric struct {
	Name       string  `json:"name"`
	QtResults  int64   `json:"qtResults"`
	TotalPrice float64 `json:"totalPrice"`
	AvgPrice   float64 `json:"avgPrice"`
}

type MetricsResponse struct {
	Metrics  []Metric `json:"metrics"`
	MinPrice float64  `json:"minPrice"`
	MaxPrice float64  `json:"maxPrice"`
}

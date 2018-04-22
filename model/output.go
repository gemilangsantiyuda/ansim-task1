package model

type Output struct {
	Policy              Policy  `json:"policy"`
	AverageTotalCost    float64 `json:"average_total_cost"`
	AverageOrderingCost float64 `json:"average_ordering_cost"`
	AverageHoldingCost  float64 `json:"average_holding_cost"`
	AverageShortageCost float64 `json:"average_shortage_cost"`
}

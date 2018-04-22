package model

type Status struct {
	InventoryLevel int     `json:"inventory_level"`
	OrderingCost   float64 `json:"ordering_cost"`
	OrderFrequency int     `json:"order_frequency"`
	HoldingCost    float64 `json:"holding_cost"`
	ShortageCost   float64 `json:"shortage_cost"`
}

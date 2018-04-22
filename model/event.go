package model

type Event struct {
	EventCode int `json:"event_code"`
	SupplyQty int `json:"supply_qty"`
	DemandQty int `json:"demand_qty"`
}

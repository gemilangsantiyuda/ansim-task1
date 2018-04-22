package model

type Input struct {
	InitialInventory      int          `json:"initial_inventory"`
	NumberDemandSizes     int          `json:"number_demand_sizes"`
	DemandSizes           []DemandSize `json:"demand_sizes"`
	MeanInterdemandTime   float64      `json:"mean_interdemand_time"`
	DeliveryLagLowerRange float64      `json:"delivery_lag_lower_range"`
	DeliveryLagUpperRange float64      `json:"delivery_lag_upper_range"`
	SimulationLength      int          `json:"simulation_length"`
	EvaluationPeriod      int          `json:"evaluation_period"`
	SetupCost             float64      `json:"setup_cost"`
	IncrementalCost       float64      `json:"incremental_cost"`
	Profit                float64      `json:"profit"`
	NumberPolicies        int          `json:"number_policies"`
	Policies              []Policy     `json:"policies"`
	DailyHoldingCost      float64      `json:"daily_holding_cost"`
}

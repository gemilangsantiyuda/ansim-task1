package model

type DemandSize struct {
	Size             int     `json:"size"`
	ProbabilityLower float64 `json:"probability_lower"`
	ProbabilityUpper float64 `json:"probability_upper"`
}

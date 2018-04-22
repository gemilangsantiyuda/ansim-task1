package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"sort"

	"github.com/campus/simulation-and-queue/task-1/model"
)

const INVENTORY_EVALUATION = 1
const DEMAND_HAPPEN = 3
const SUPPLY_ARRIVE = 2
const MONTH_IN_DAYS = 30

func main() {
	var input model.Input
	file, err := ioutil.ReadFile("input.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &input)
	if err != nil {
		panic(err)
	}

	reportList, err := simulate(input)
	if err != nil {
		panic(err)
	}

	report, err := json.Marshal(reportList)
	ioutil.WriteFile("output.json", report, 0644)
}

func simulate(input model.Input) ([]model.Output, error) {

	var reportList []model.Output
	for policyIdx := range input.Policies {
		policyReport, err := getPolicyReport(input, policyIdx)
		if err != nil {
			panic(err)
		}
		reportList = append(reportList, policyReport)
	}

	return reportList, nil
}

func getPolicyReport(input model.Input, policyIdx int) (model.Output, error) {

	/*INITIATION*/
	policy := input.Policies[policyIdx]
	totalDays := input.SimulationLength * MONTH_IN_DAYS
	var eventsList [][]model.Event
	for day := 0; day < totalDays; day++ {
		var events []model.Event
		eventsList = append(eventsList, events)
	}

	evaluationDay := 0
	for evaluationDay < totalDays {
		eventsList[evaluationDay] = append(eventsList[evaluationDay], model.Event{
			INVENTORY_EVALUATION,
			0,
			0,
		})
		evaluationDay += input.EvaluationPeriod * MONTH_IN_DAYS
	}

	demandDay := 0
	for demandDay < totalDays {
		nextDemandDay := demandDay + int(rand.ExpFloat64()/10*MONTH_IN_DAYS)
		if nextDemandDay == demandDay {
			continue
		}
		demandDay = nextDemandDay
		if demandDay >= totalDays {
			break
		}
		qtyProbability := float64(rand.Intn(1000)) / 1000.000000
		qty := 0
		for _, demandSize := range input.DemandSizes {
			if qtyProbability >= demandSize.ProbabilityLower && qtyProbability < demandSize.ProbabilityUpper {
				qty = demandSize.Size
				break
			}
		}
		eventsList[demandDay] = append(eventsList[demandDay], model.Event{
			DEMAND_HAPPEN,
			0,
			qty,
		})
	}

	var status model.Status
	/*INITIATION*/

	/*SIMULATE EACH DAY*/
	for today := 0; today < totalDays; today++ {
		events := eventsList[today]
		sort.SliceStable(events, func(i int, j int) bool {
			return events[i].EventCode < events[j].EventCode
		})
		for _, event := range events {
			switch event.EventCode {
			case INVENTORY_EVALUATION:
				if status.InventoryLevel < policy.Lower {
					supplyOrderQty := policy.Upper - status.InventoryLevel
					supplyLag := int(float64(rand.Intn(100)+100) / 200.000 * MONTH_IN_DAYS)
					supplyArrivalDay := today + supplyLag
					eventsList[supplyArrivalDay] = append(eventsList[supplyArrivalDay], model.Event{
						SUPPLY_ARRIVE,
						supplyOrderQty,
						0,
					})
					status.OrderingCost += input.SetupCost + input.IncrementalCost*float64(supplyOrderQty)
					status.OrderFrequency++
				}

			case SUPPLY_ARRIVE:
				status.InventoryLevel += event.SupplyQty
			case DEMAND_HAPPEN:
				if event.DemandQty <= status.InventoryLevel {
					status.InventoryLevel -= event.DemandQty
				} else {
					status.ShortageCost += float64(event.DemandQty) * input.Profit
				}
			}
		}
		status.HoldingCost += float64(status.InventoryLevel) * input.DailyHoldingCost
	}
	averageOrderingCost := status.OrderingCost / float64(input.SimulationLength)
	averageHoldingCost := status.HoldingCost / float64(input.SimulationLength)
	averageShortageCost := status.ShortageCost / float64(input.SimulationLength)
	averageTotalCost := averageOrderingCost + averageHoldingCost + averageShortageCost
	return model.Output{
		policy,
		averageTotalCost,
		averageOrderingCost,
		averageHoldingCost,
		averageShortageCost,
	}, nil
}

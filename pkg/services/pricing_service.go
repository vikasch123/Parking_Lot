package services

import (
	"parking-lot/pkg/vehicle"
	"time"
)

type PricingService struct {
	calculators map[string]Payable
}

func NewPricingService() *PricingService {
	return &PricingService{
		calculators: map[string]Payable{
			"Small":    &SmallVehicleCalculator{},
			"Large":    &LargeVehicleCalculator{},
			"Handicap": &HandicapVehicleCalculator{},
		},
	}
}

func (s *PricingService) CalculateCost(v vehicle.Vehicle, parkedAt time.Time) float64 {
	duration := time.Since(parkedAt)

	var key string
	if v.IsHandicap {
		key = "Handicap"
	} else {
		key = v.Size // "Small" or "Large"
	}

	calculator, ok := s.calculators[key]
	if !ok {
		calculator = &SmallVehicleCalculator{} // default
	}

	return calculator.CalculateAmountPayable(v, duration)
}

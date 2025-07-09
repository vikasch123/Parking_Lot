package services

import (
	// "parking-lot/pkg/attendant"
	// "parking-lot/pkg/lot"
	"parking-lot/pkg/vehicle"
	"time"
)

type Payable interface {
	CalculateAmountPayable(v vehicle.Vehicle, t time.Duration) float64
}

type SmallVehicleCalculator struct{}

func (s *SmallVehicleCalculator) CalculateAmountPayable(v vehicle.Vehicle, t time.Duration) float64 {
	hours := t.Hours()
	return 30 * hours
}

type LargeVehicleCalculator struct{}

func (c *LargeVehicleCalculator) CalculateAmountPayable(v vehicle.Vehicle, duration time.Duration) float64 {
	hours := duration.Hours()
	return 50.0 * hours // ₹50/hour
}

type HandicapVehicleCalculator struct{}

func (c *HandicapVehicleCalculator) CalculateAmountPayable(v vehicle.Vehicle, duration time.Duration) float64 {
	hours := duration.Hours()
	return 5.0 * hours // discounted
}

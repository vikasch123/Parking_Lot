package services

import (
	"testing"
	"time"

	"parking-lot/pkg/vehicle"

	"github.com/stretchr/testify/assert"
)

func TestPricingService_CalculateCost_Small(t *testing.T) {
	service := NewPricingService()

	car := vehicle.Vehicle{
		Plate: "DL-1234", Color: "White", Make: "Suzuki",
		Size: "Small", IsHandicap: false,
	}

	parkedAt := time.Now().Add(-2 * time.Hour) // parked 2 hours ago.. a
	cost := service.CalculateCost(car, parkedAt)

	expected := 30.0 * 2
	assert.InEpsilon(t, expected, cost, 0.01)
}

func TestPricingService_CalculateCost_Large(t *testing.T) {
	service := NewPricingService()

	car := vehicle.Vehicle{
		Plate: "DL-5678", Color: "Black", Make: "Fortuner",
		Size: "Large", IsHandicap: false,
	}

	parkedAt := time.Now().Add(-1 * time.Hour) // 1 hour
	cost := service.CalculateCost(car, parkedAt)

	expected := 50.0 * 1
	assert.InEpsilon(t, expected, cost, 0.01)
}

func TestPricingService_CalculateCost_Handicap(t *testing.T) {
	service := NewPricingService()

	car := vehicle.Vehicle{
		Plate: "DL-9999", Color: "Blue", Make: "WagonR",
		Size: "Small", IsHandicap: true,
	}

	parkedAt := time.Now().Add(-3 * time.Hour)
	cost := service.CalculateCost(car, parkedAt)

	expected := 5.0 * 3
	assert.InEpsilon(t, expected, cost, 0.01)
}

func TestPricingService_CalculateCost_UnknownType_DefaultsToSmall(t *testing.T) {
	service := NewPricingService()

	car := vehicle.Vehicle{
		Plate: "XX-0000", Color: "Grey", Make: "TestCar",
		Size: "Unknown", IsHandicap: false,
	}

	parkedAt := time.Now().Add(-1 * time.Hour)
	cost := service.CalculateCost(car, parkedAt)

	expected := 30.0 // fallback to small vehicle rate
	assert.InEpsilon(t, expected, cost, 0.01)
}

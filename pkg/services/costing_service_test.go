package services

import (
	"testing"
	"time"

	"parking-lot/pkg/vehicle"

	"github.com/stretchr/testify/assert"
)

func TestSmallVehicleCalculator(t *testing.T) {
	calculator := &SmallVehicleCalculator{}
	car := vehicle.Vehicle{
		Plate: "DL-1234", Color: "White", Make: "Suzuki", Size: "Small", IsHandicap: false,
	}

	duration := 2 * time.Hour
	expected := 60.0 // ₹30/hour * 2 hours

	result := calculator.CalculateAmountPayable(car, duration)
	assert.Equal(t, expected, result)
}

func TestLargeVehicleCalculator(t *testing.T) {
	calculator := &LargeVehicleCalculator{}
	vehicle := vehicle.Vehicle{
		Plate: "DL-5678", Color: "Black", Make: "Fortuner", Size: "Large", IsHandicap: false,
	}

	duration := 90 * time.Minute
	expected := 75.0 // ₹50/hour * 1.5 hours

	result := calculator.CalculateAmountPayable(vehicle, duration)
	assert.Equal(t, expected, result)
}

func TestHandicapVehicleCalculator(t *testing.T) {
	calculator := &HandicapVehicleCalculator{}
	vehicle := vehicle.Vehicle{
		Plate: "DL-9999", Color: "Blue", Make: "WagonR", Size: "Small", IsHandicap: true,
	}

	duration := 3 * time.Hour
	expected := 15.0 // ₹5/hour * 3 hours

	result := calculator.CalculateAmountPayable(vehicle, duration)
	assert.Equal(t, expected, result)
}

package services

import (
	"testing"
	"parking-lot/pkg/attendant"
	"parking-lot/pkg/lot"
	"parking-lot/pkg/vehicle"
	"github.com/stretchr/testify/assert"
)

func setupTestPoliceService() *PoliceService {
	lot1 := lot.NewParkingLot("Lot A", 3, make(map[string]lot.ParkedVehicle))
	_, _ = lot1.Park(*vehicle.New("DL-1001", "White", "BMW", false, "Small"))
	_, _ = lot1.Park(*vehicle.New("DL-1002", "Black", "Audi", false, "Small"))
	att := attendant.NewAttendant("Ravi", []*lot.ParkingLot{lot1})
	return NewPoliceService(att)
}

func TestFindCarByColor(t *testing.T) {
	service := setupTestPoliceService()
	vehicles, _ := service.FindCarByColor("white")
	assert.Len(t, vehicles, 1)
	assert.Equal(t, "DL-1001", vehicles[0].Plate)
}

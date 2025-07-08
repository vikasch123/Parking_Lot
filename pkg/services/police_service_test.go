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

func TestGetVehiclePlatesByColor(t *testing.T) {
	service := setupTestPoliceService()
	plates := service.GetVehiclePlatesByColor("white")
	assert.Len(t, plates, 1)
	assert.Contains(t, plates, "DL-1001")
}

func TestGetLotByVehicleNumber(t *testing.T) {
	service := setupTestPoliceService()
	lotName, found := service.GetLotByVehicleNumber("DL-1001")
	assert.True(t, found)
	assert.Equal(t, "Lot A", lotName)
	_, found = service.GetLotByVehicleNumber("DL-9999")
	assert.False(t, found)
}

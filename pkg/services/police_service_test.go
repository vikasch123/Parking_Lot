package services

import (
	"parking-lot/pkg/attendant"
	"parking-lot/pkg/lot"
	"parking-lot/pkg/vehicle"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setupTestPoliceService() *PoliceService {
	lot1 := lot.NewParkingLot("Lot A", 3, make(map[string]lot.ParkedVehicle))
	lot2 := lot.NewParkingLot("Lot B", 3, make(map[string]lot.ParkedVehicle))

	// Park vehicles...
	_, _ = lot1.Park(*vehicle.New("DL-1001", "White", "BMW", false, "Small"))
	time.Sleep(1 * time.Second)
	_, _ = lot1.Park(*vehicle.New("DL-1002", "Black", "Audi", false, "Small"))
	time.Sleep(1 * time.Second)
	_, _ = lot2.Park(*vehicle.New("DL-1003", "WHITE", "WagonR", true, "Small"))
	time.Sleep(1 * time.Second)
	_, _ = lot2.Park(*vehicle.New("DL-1004", "Red", "Fortuner", false, "Large"))

	att := attendant.NewAttendant("Ravi", []*lot.ParkingLot{lot1, lot2})
	return NewPoliceService(att)
}

func TestFindCarByColor(t *testing.T) {
	service := setupTestPoliceService()
	vehicles, _ := service.FindCarByColor("white")

	assert.Len(t, vehicles, 2)
	assert.Equal(t, "DL-1001", vehicles[0].Plate)
	assert.Equal(t, "DL-1003", vehicles[1].Plate)
}

func TestGetVehiclePlatesByColor(t *testing.T) {
	service := setupTestPoliceService()
	plates := service.GetVehiclePlatesByColor("white")

	assert.Len(t, plates, 2)
	assert.Contains(t, plates, "DL-1001")
	assert.Contains(t, plates, "DL-1003")
}

func TestGetLotByVehicleNumber(t *testing.T) {
	service := setupTestPoliceService()

	lotName, found := service.GetLotByVehicleNumber("DL-1003")
	assert.True(t, found)
	assert.Equal(t, "Lot B", lotName)

	_, found = service.GetLotByVehicleNumber("DL-9999")
	assert.False(t, found)
}

func TestGetVehiclesParkedAfter(t *testing.T) {
	service := setupTestPoliceService()

	after := time.Now().Add(-2 * time.Second)
	vehicles := service.GetVehiclesParkedAfter(after)

	assert.GreaterOrEqual(t, len(vehicles), 1)
}

func TestGetAllParkedVehiclesSummary(t *testing.T) {
	service := setupTestPoliceService()

	summary := service.GetAllParkedVehiclesSummary()

	assert.Len(t, summary, 4)
	plates := map[string]bool{}
	for _, entry := range summary {
		plates[entry["Plate"]] = true
	}
	assert.True(t, plates["DL-1001"])
	assert.True(t, plates["DL-1002"])
	assert.True(t, plates["DL-1003"])
	assert.True(t, plates["DL-1004"])
}

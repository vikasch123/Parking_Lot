package services

import (
	"os"
	"testing"
	"time"

	"parking-lot/pkg/attendant"
	"parking-lot/pkg/lot"
	"parking-lot/pkg/vehicle"

	"github.com/stretchr/testify/assert"
)

func setupExportTest() *PoliceService {
	lot1 := lot.NewParkingLot("Lot A", 3, make(map[string]lot.ParkedVehicle))
	lot2 := lot.NewParkingLot("Lot B", 3, make(map[string]lot.ParkedVehicle))

	_, _ = lot1.Park(*vehicle.New("DL-1001", "White", "BMW", false, "Small"))
	_, _ = lot2.Park(*vehicle.New("DL-1002", "White", "Honda", false, "Small"))
	_, _ = lot2.Park(*vehicle.New("DL-1003", "Red", "Kia", false, "Large"))

	att := attendant.NewAttendant("Ravi", []*lot.ParkingLot{lot1, lot2})
	return NewPoliceService(att)
}

func TestExportVehiclesByColorToCSV(t *testing.T) {
	service := setupExportTest()
	filename := "test_out/vehicles_white.csv"

	_ = os.MkdirAll("test_out", os.ModePerm)
	vehicles := service.ExportVehiclesByColorToCSV("White", filename)

	assert.Len(t, vehicles, 2)

	// Check if file exists
	_, err := os.Stat(filename)
	assert.NoError(t, err)
}

func TestExportPlatesByColorToCSV(t *testing.T) {
	service := setupExportTest()
	filename := "test_out/plates_white.csv"

	err := service.ExportPlatesByColorToCSV("White", filename)
	assert.Nil(t, err)

	_, err = os.Stat(filename)
	assert.NoError(t, err)
}

func TestExportVehiclesParkedAfterToCSV(t *testing.T) {
	service := setupExportTest()
	filename := "test_out/after.csv"
	after := time.Now().Add(-2 * time.Hour)

	err := service.ExportVehiclesParkedAfterToCSV(after, filename)
	assert.Nil(t, err)

	_, err = os.Stat(filename)
	assert.NoError(t, err)
}

func TestExportAllParkedVehiclesSummaryToCSV(t *testing.T) {
	service := setupExportTest()
	filename := "test_out/summary.csv"

	err := service.ExportAllParkedVehiclesSummaryToCSV(filename)
	assert.Nil(t, err)

	_, err = os.Stat(filename)
	assert.NoError(t, err)
}

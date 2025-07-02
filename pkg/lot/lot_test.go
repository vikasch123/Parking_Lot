package lot

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParkingLotStruct(t *testing.T) {
	lot := ParkingLot{
		Name:     "Lot A",
		Capacity: 10,
		Vehicles: make(map[string]ParkedVehicle),
	}
	assert.Equal(t, "Lot A", lot.Name)
	assert.Equal(t, 10, lot.Capacity)
	assert.NotNil(t, lot.Vehicles)
}

func TestNewParkingLot(t *testing.T) {
	lot := NewParkingLot("Lot B", 5, make(map[string]ParkedVehicle))
	assert.NotNil(t, lot)
	assert.Equal(t, "Lot B", lot.Name)
	assert.Equal(t, 5, lot.Capacity)
	assert.NotNil(t, lot.Vehicles)
}

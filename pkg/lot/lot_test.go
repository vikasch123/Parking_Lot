package lot

import (
	"parking-lot/pkg/vehicle"
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

func TestParkCar_Success(t *testing.T) {
	lot := NewParkingLot("B2", 2, make(map[string]ParkedVehicle))
	car := vehicle.New("RJ-14-AB-6567", "White", "Wagonr", false, "small")
	ok, err := lot.Park(*car)
	assert.True(t, ok)
	assert.Nil(t, err)
}

func TestUnparkCar_Success(t *testing.T) {
	lot := NewParkingLot("Lot A", 2, make(map[string]ParkedVehicle))
	car := vehicle.New("DL-1234", "White", "BMW", false, "Small")
	_, _ = lot.Park(*car)
	unparkedCar, err := lot.Unpark("DL-1234")
	assert.Nil(t, err)
	assert.Equal(t, *car, unparkedCar)
}

func TestIsFullAndIsAvailable(t *testing.T) {
	lot := NewParkingLot("Lot A", 1, make(map[string]ParkedVehicle))
	car := vehicle.New("DL-1234", "White", "BMW", false, "Small")

	assert.True(t, lot.IsAvailable())
	assert.False(t, lot.IsFull())

	_, _ = lot.Park(*car)

	assert.False(t, lot.IsAvailable())
	assert.True(t, lot.IsFull())
}

func TestFreeSlots(t *testing.T) {
	lot := NewParkingLot("Lot A", 2, make(map[string]ParkedVehicle))
	assert.Equal(t, 2, lot.FreeSlots())
	car := vehicle.New("DL-1234", "White", "BMW", false, "Small")
	_, _ = lot.Park(*car)
	assert.Equal(t, 1, lot.FreeSlots())
}

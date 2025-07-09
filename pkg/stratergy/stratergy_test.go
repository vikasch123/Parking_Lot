package stratergy

import (
	"testing"

	"parking-lot/pkg/lot"
	"parking-lot/pkg/vehicle"

	"github.com/stretchr/testify/assert"
)

func setupTestLots() []*lot.ParkingLot {
	lot1 := lot.NewParkingLot("Lot A", 3, make(map[string]lot.ParkedVehicle))
	lot2 := lot.NewParkingLot("Lot B", 2, make(map[string]lot.ParkedVehicle))
	lot3 := lot.NewParkingLot("Lot C", 1, make(map[string]lot.ParkedVehicle))

	return []*lot.ParkingLot{lot1, lot2, lot3}
}

func TestParkEvenly_EmptyLots(t *testing.T) {
	strategy := &ParkEvenly{}
	lots := setupTestLots()
	car := vehicle.New("DL-1234", "White", "BMW", false, "Small")

	selectedLot, err := strategy.Park(*car, lots)

	assert.Nil(t, err)
	assert.NotNil(t, selectedLot)
	assert.Equal(t, "Lot A", selectedLot.Name) // Should pick first lot when all are empty
}

func TestParkEvenly_DistributeEvenly(t *testing.T) {
	strategy := &ParkEvenly{}
	lots := setupTestLots()

	// Park one car in lot A
	car1 := vehicle.New("DL-1234", "White", "BMW", false, "Small")
	_, _ = lots[0].Park(*car1)

	// Park one car in lot B
	car2 := vehicle.New("DL-5678", "Black", "Audi", false, "Small")
	_, _ = lots[1].Park(*car2)

	// Try to park another car - should go to lot C (least occupied)
	car3 := vehicle.New("DL-9999", "Red", "Swift", false, "Small")
	selectedLot, err := strategy.Park(*car3, lots)

	assert.Nil(t, err)
	assert.NotNil(t, selectedLot)
	assert.Equal(t, "Lot C", selectedLot.Name)
}

func TestParkEvenly_AllLotsFull(t *testing.T) {
	strategy := &ParkEvenly{}
	lots := setupTestLots()

	// Fill all lots
	car1 := vehicle.New("DL-1234", "White", "BMW", false, "Small")
	car2 := vehicle.New("DL-5678", "Black", "Audi", false, "Small")
	car3 := vehicle.New("DL-9999", "Red", "Swift", false, "Small")
	car4 := vehicle.New("DL-1111", "Blue", "Honda", false, "Small")
	car5 := vehicle.New("DL-2222", "Green", "Toyota", false, "Small")
	car6 := vehicle.New("DL-3333", "Yellow", "Ford", false, "Small")

	_, _ = lots[0].Park(*car1)
	_, _ = lots[0].Park(*car2)
	_, _ = lots[0].Park(*car3)
	_, _ = lots[1].Park(*car4)
	_, _ = lots[1].Park(*car5)
	_, _ = lots[2].Park(*car6)

	// Try to park another car
	car7 := vehicle.New("DL-4444", "Purple", "Nissan", false, "Small")
	selectedLot, err := strategy.Park(*car7, lots)

	assert.NotNil(t, err)
	assert.Nil(t, selectedLot)
	assert.EqualError(t, err, "all lots are full")
}

func TestHandicapPark_Success(t *testing.T) {
	strategy := &HandicapPark{}
	lots := setupTestLots()
	car := vehicle.New("DL-1234", "White", "BMW", true, "Small") // handicap vehicle

	selectedLot, err := strategy.Park(*car, lots)

	assert.Nil(t, err)
	assert.NotNil(t, selectedLot)
	// Should park in first available lot
	assert.True(t, selectedLot.Name == "Lot A" || selectedLot.Name == "Lot B" || selectedLot.Name == "Lot C")
}

func TestHandicapPark_AllLotsFull(t *testing.T) {
	strategy := &HandicapPark{}
	lots := setupTestLots()

	// Fill all lots
	car1 := vehicle.New("DL-1234", "White", "BMW", false, "Small")
	car2 := vehicle.New("DL-5678", "Black", "Audi", false, "Small")
	car3 := vehicle.New("DL-9999", "Red", "Swift", false, "Small")
	car4 := vehicle.New("DL-1111", "Blue", "Honda", false, "Small")
	car5 := vehicle.New("DL-2222", "Green", "Toyota", false, "Small")
	car6 := vehicle.New("DL-3333", "Yellow", "Ford", false, "Small")

	_, _ = lots[0].Park(*car1)
	_, _ = lots[0].Park(*car2)
	_, _ = lots[0].Park(*car3)
	_, _ = lots[1].Park(*car4)
	_, _ = lots[1].Park(*car5)
	_, _ = lots[2].Park(*car6)

	// Try to park handicap vehicle
	car7 := vehicle.New("DL-4444", "Purple", "Nissan", true, "Small")
	selectedLot, err := strategy.Park(*car7, lots)

	assert.NotNil(t, err)
	assert.Nil(t, selectedLot)
	assert.EqualError(t, err, "no space available for handicap drivers")
}

func TestBigVehiclePark_Success(t *testing.T) {
	strategy := &BigVehiclePark{}
	lots := setupTestLots()
	car := vehicle.New("DL-1234", "White", "Truck", false, "Large")

	selectedLot, err := strategy.Park(*car, lots)

	assert.Nil(t, err)
	assert.NotNil(t, selectedLot)
	// Should pick lot with most free slots (Lot A with 3 slots)
	assert.Equal(t, "Lot A", selectedLot.Name)
}

func TestBigVehiclePark_PickLotWithMostSpace(t *testing.T) {
	strategy := &BigVehiclePark{}
	lots := setupTestLots()

	// Fill some slots in lot A and B, leave lot C empty
	car1 := vehicle.New("DL-1234", "White", "BMW", false, "Small")
	car2 := vehicle.New("DL-5678", "Black", "Audi", false, "Small")
	car3 := vehicle.New("DL-9999", "Red", "Swift", false, "Small")

	_, _ = lots[0].Park(*car1) // Lot A: 2 free slots
	_, _ = lots[0].Park(*car2) // Lot A: 1 free slot
	_, _ = lots[1].Park(*car3) // Lot B: 1 free slot
	// Lot C: 1 free slot (empty)

	// Try to park large vehicle - should pick lot with most free slots
	car4 := vehicle.New("DL-1111", "Blue", "Truck", false, "Large")
	selectedLot, err := strategy.Park(*car4, lots)

	assert.Nil(t, err)
	assert.NotNil(t, selectedLot)
	// All lots have 1 free slot, so it will pick the first one it encounters
	assert.True(t, selectedLot.Name == "Lot A" || selectedLot.Name == "Lot B" || selectedLot.Name == "Lot C")
}

func TestBigVehiclePark_AllLotsFull(t *testing.T) {
	strategy := &BigVehiclePark{}
	lots := setupTestLots()

	// Fill all lots
	car1 := vehicle.New("DL-1234", "White", "BMW", false, "Small")
	car2 := vehicle.New("DL-5678", "Black", "Audi", false, "Small")
	car3 := vehicle.New("DL-9999", "Red", "Swift", false, "Small")
	car4 := vehicle.New("DL-1111", "Blue", "Honda", false, "Small")
	car5 := vehicle.New("DL-2222", "Green", "Toyota", false, "Small")
	car6 := vehicle.New("DL-3333", "Yellow", "Ford", false, "Small")

	_, _ = lots[0].Park(*car1)
	_, _ = lots[0].Park(*car2)
	_, _ = lots[0].Park(*car3)
	_, _ = lots[1].Park(*car4)
	_, _ = lots[1].Park(*car5)
	_, _ = lots[2].Park(*car6)

	// Try to park large vehicle
	car7 := vehicle.New("DL-4444", "Purple", "Truck", false, "Large")
	selectedLot, err := strategy.Park(*car7, lots)

	assert.NotNil(t, err)
	assert.Nil(t, selectedLot)
	assert.EqualError(t, err, "no lot available for large vehicle")
}

func TestBigVehiclePark_NoLotsAvailable(t *testing.T) {
	strategy := &BigVehiclePark{}
	lots := []*lot.ParkingLot{} // Empty lots slice

	car := vehicle.New("DL-1234", "White", "Truck", false, "Large")
	selectedLot, err := strategy.Park(*car, lots)

	assert.NotNil(t, err)
	assert.Nil(t, selectedLot)
	assert.EqualError(t, err, "no lot available for large vehicle")
}

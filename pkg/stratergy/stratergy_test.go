
package stratergy

import (
	"testing"

	"parking-lot/pkg/lot"
	"parking-lot/pkg/vehicle"

	"github.com/stretchr/testify/assert"
)

func TestHandicapPark_Success(t *testing.T) {
	strategy := &HandicapPark{}
	lots := []*lot.ParkingLot{
		lot.NewParkingLot("Lot A", 2, make(map[string]lot.ParkedVehicle)),
	}
	car := vehicle.New("DL-5678", "White", "WagonR", true, "Small")
	selectedLot, err := strategy.Park(*car, lots)
	assert.Nil(t, err)
	assert.NotNil(t, selectedLot)
}

package attendant

import (
	"testing"
	"parking-lot/pkg/lot"
	"github.com/stretchr/testify/assert"
)

func TestNewAttendant(t *testing.T) {
	lots := []*lot.ParkingLot{
		lot.NewParkingLot("Lot A", 2, make(map[string]lot.ParkedVehicle)),
	}
	att := NewAttendant("Ravi", lots)
	assert.NotNil(t, att)
	assert.Equal(t, "Ravi", att.Name)
	assert.Equal(t, lots, att.Lots)
}

package vehicle

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestVehicleStruct(t *testing.T) {
	v := Vehicle{
		Plate:      "DL-1234",
		Color:      "White",
		Make:       "BMW",
		IsHandicap: false,
		Size:       "Small",
	}
	assert.Equal(t, "DL-1234", v.Plate)
	assert.Equal(t, "White", v.Color)
	assert.Equal(t, "BMW", v.Make)
	assert.False(t, v.IsHandicap)
	assert.Equal(t, "Small", v.Size)
}

func TestNewVehicle(t *testing.T) {
	car := New("DL-1234", "White", "BMW", false, "Small")
	assert.NotNil(t, car)
	assert.Equal(t, "DL-1234", car.Plate)
	assert.Equal(t, "White", car.Color)
	assert.Equal(t, "BMW", car.Make)
	assert.False(t, car.IsHandicap)
	assert.Equal(t, "Small", car.Size)
}

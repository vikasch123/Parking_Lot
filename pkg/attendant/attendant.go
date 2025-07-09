package attendant

import (
	"parking-lot/pkg/lot"
	"parking-lot/pkg/stratergy"
	"parking-lot/pkg/vehicle"
)

type Attendant struct {
	Name string
	Lots []*lot.ParkingLot
}

func NewAttendant(name string, lots []*lot.ParkingLot) *Attendant {
	return &Attendant{
		Name: name,
		Lots: lots,
	}
}

// ParkEvenly tries to park a vehicle using even distribution
// Generic park method using a strategy..
// ParkingStratergy - interface in stratergy.go
// UC10   - attendant select strategy dynamically
func (a *Attendant) ParkWithStrategy(v vehicle.Vehicle, s stratergy.ParkingStratergy) (*lot.ParkingLot, error) {
	return s.Park(v, a.Lots)
}

package attendant

import "parking-lot/pkg/lot"

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

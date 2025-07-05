
package stratergy

import (
	"errors"
	"parking-lot/pkg/lot"
	"parking-lot/pkg/vehicle"
)


type HandicapPark struct{}

func (h *HandicapPark) Park(v vehicle.Vehicle, lots []*lot.ParkingLot) (*lot.ParkingLot, error) {
	for _, l := range lots {
		if !l.IsFull() {
			_, err := l.Park(v)
			if err != nil {
				return nil, err
			}
			return l, nil
		}
	}
	return nil, errors.New("no space available for handicap drivers")
}

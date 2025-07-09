package stratergy

import (
	"errors"
	"parking-lot/pkg/lot"
	"parking-lot/pkg/vehicle"
)

// This is the interface for generic parking stratergy
type ParkingStratergy interface {
	Park(v vehicle.Vehicle, lots []*lot.ParkingLot) (*lot.ParkingLot, error)
}

type ParkEvenly struct{}

func (p *ParkEvenly) Park(v vehicle.Vehicle, lots []*lot.ParkingLot) (*lot.ParkingLot, error) {
	var targetLot *lot.ParkingLot
	minCars := int(^uint(0) >> 1)

	for _, l := range lots {
		if l.IsFull() {
			continue
		}
		if len(l.Vehicles) < minCars {
			minCars = len(l.Vehicles)
			targetLot = l
		}
	}

	if targetLot == nil {
		return nil, errors.New("all lots are full")
	}

	_, err := targetLot.Park(v)
	if err != nil {
		return nil, err
	}

	return targetLot, nil
}

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

type BigVehiclePark struct{}

// Implements UC9.
func (b *BigVehiclePark) Park(v vehicle.Vehicle, lots []*lot.ParkingLot) (*lot.ParkingLot, error) {
	var targetLot *lot.ParkingLot
	maxfree := -1
	for _, l := range lots {
		free := l.FreeSlots()
		if free > maxfree {
			maxfree = free
			targetLot = l
		}
	}

	if targetLot == nil || maxfree == 0 {
		return nil, errors.New("no lot available for large vehicle")
	}

	_, err := targetLot.Park(v)
	if err != nil {
		return nil, err
	}
	return targetLot, nil
}

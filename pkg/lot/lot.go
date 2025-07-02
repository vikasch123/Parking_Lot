package lot

import "parking-lot/pkg/vehicle"

type ParkedVehicle struct {
	Vehicle vehicle.Vehicle
}

type ParkingLot struct {
	Name     string
	Capacity int
	Vehicles map[string]ParkedVehicle
}

func NewParkingLot(name string, cap int, vehicles map[string]ParkedVehicle) *ParkingLot {
	return &ParkingLot{
		Name:     name,
		Capacity: cap,
		Vehicles: vehicles,
	}
}

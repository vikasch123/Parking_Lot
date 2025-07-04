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

func (p *ParkingLot) Park(v vehicle.Vehicle) (bool,error){
	if len(p.Vehicles) >= p.Capacity {
		return false, errors.New("Sorry,Parking lot is full")
	}
	if _, exists := p.Vehicles[v.Plate]; exists {
		return false, errors.New("trying to park already parked car")
	}
	p.Vehicles[v.Plate] = ParkedVehicle{
		Vehicle:  v,
		ParkedAt: time.Now(),
	}
	return true, nil
}

func (p *ParkingLot) Unpark(plate string) (vehicle.Vehicle, error) {
	parked, exists := p.Vehicles[plate]
	if !exists {
		return vehicle.Vehicle{}, errors.New("vehicle not found")
	}
	delete(p.Vehicles, plate)
	return parked.Vehicle, nil
}

func (p *ParkingLot) IsFull() bool {
	return len(p.Vehicles) >= p.Capacity
}

func (p *ParkingLot) IsAvailable() bool {
	return len(p.Vehicles) < p.Capacity
}

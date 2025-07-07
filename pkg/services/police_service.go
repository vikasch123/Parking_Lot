package services

import (
	"parking-lot/pkg/attendant"
	"parking-lot/pkg/vehicle"
	"strings"
)

type PoliceService struct {
	attendant *attendant.Attendant
}

func NewPoliceService(att *attendant.Attendant) *PoliceService {
	return &PoliceService{
		attendant: att,
	}
}

// UC13: Find cars by color
func (p *PoliceService) FindCarByColor(color string) ([]vehicle.Vehicle, error) {
	var result []vehicle.Vehicle
	for _, l := range p.attendant.Lots {
		for _, pv := range l.GetParkedVehicles() {
			if strings.EqualFold(pv.Vehicle.Color, color) {
				result = append(result, pv.Vehicle)
			}
		}
	}
	return result, nil
}

func (p *PoliceService) GetVehiclePlatesByColor(color string) []string {
	var result []string
	for _, lot := range p.attendant.Lots {
		for _, pv := range lot.GetParkedVehicles() {
			if strings.EqualFold(pv.Vehicle.Color, color) {
				result = append(result, pv.Vehicle.Plate)
			}
		}
	}
	return result
}

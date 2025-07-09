package services

import (
	"fmt"
	"parking-lot/pkg/attendant"
	"parking-lot/pkg/vehicle"
	"strings"
	"time"
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

// UC15: Find lot by vehicle number
func (p *PoliceService) GetLotByVehicleNumber(number string) (string, bool) {
	for _, v := range p.attendant.Lots {
		for _, pv := range v.GetParkedVehicles() {
			if strings.EqualFold(pv.Vehicle.Plate, number) {
				return v.Name, true
			}
		}
	}
	return "", false
}

// UC16: Get all vehicles parked after a certain time
func (p *PoliceService) GetVehiclesParkedAfter(after time.Time) []vehicle.Vehicle {
	var result []vehicle.Vehicle
	for _, lot := range p.attendant.Lots {
		for _, pv := range lot.GetParkedVehicles() {
			if pv.ParkedAt.After(after) {
				result = append(result, pv.Vehicle)
			}
		}
	}
	return result
}

// UC 17 As a police officer, I want to see a summary of all parked vehicles in all lots.
func (p *PoliceService) GetAllParkedVehiclesSummary() []map[string]string {
	var result []map[string]string

	for _, lot := range p.attendant.Lots {
		for _, pv := range lot.GetParkedVehicles() {
			entry := map[string]string{
				"Lot":        lot.Name,
				"Plate":      pv.Vehicle.Plate,
				"Make":       pv.Vehicle.Make,
				"Color":      pv.Vehicle.Color,
				"Size":       pv.Vehicle.Size,
				"IsHandicap": fmt.Sprintf("%v", pv.Vehicle.IsHandicap),
				"ParkedAt":   pv.ParkedAt.Format("2006-01-02 15:04:05"),
			}
			result = append(result, entry)
		}
	}

	return result
}

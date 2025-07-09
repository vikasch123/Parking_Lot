//This file just exports all the fields queried by police service
//Uses the pkg/utils/csv_export.go file for writing the output as a csv file

package services

import (
	"fmt"
	"parking-lot/pkg/utils"
	"parking-lot/pkg/vehicle"
	"time"
)

// ExportPlatesByColorToCSV fetches plates and exports them to a CSV

func (p *PoliceService) ExportVehiclesByColorToCSV(color, filename string) []vehicle.Vehicle {
	vehicles, _ := p.FindCarByColor(color)

	if len(vehicles) == 0 {
		return nil
	}
	var rows [][]string

	for _, v := range vehicles {
		row := []string{
			v.Plate,
			v.Make,
			v.Color,
			v.Size,
			fmt.Sprintf("%v", v.IsHandicap),
		}
		rows = append(rows, row)
	}

	header := []string{"Plate", "Make", "Color", "Size", "IsHandicap"}
	err := utils.ExportToCSV(header, rows, filename)

	if err != nil {
		return nil
	}

	return vehicles

}

func (p *PoliceService) ExportPlatesByColorToCSV(color, filename string) error {
	plates := p.GetVehiclePlatesByColor(color)

	if len(plates) == 0 {
		return fmt.Errorf("no vehicles found with color: %s", color)
	}

	var rows [][]string
	for _, plate := range plates {
		rows = append(rows, []string{plate})
	}

	headers := []string{"Plate Number"}
	return utils.ExportToCSV(headers, rows, filename)
}

func (p *PoliceService) ExportVehiclesParkedAfterToCSV(after time.Time, filename string) error {
	vehicles := p.GetVehiclesParkedAfter(after)

	if len(vehicles) == 0 {
		return fmt.Errorf("no vehicles parked after %v", after)
	}

	var rows [][]string
	for _, v := range vehicles {
		rows = append(rows, []string{
			v.Plate,
			v.Make,
			v.Color,
			v.Size,
			fmt.Sprintf("%v", v.IsHandicap),
		})
	}

	headers := []string{"Plate", "Make", "Color", "Size", "IsHandicap"}
	return utils.ExportToCSV(headers, rows, filename)
}

func (p *PoliceService) ExportAllParkedVehiclesSummaryToCSV(filename string) error {
	data := p.GetAllParkedVehiclesSummary()

	if len(data) == 0 {
		return fmt.Errorf("no vehicles parked in any lot")
	}

	headers := []string{"Lot", "Plate", "Make", "Color", "Size", "IsHandicap", "ParkedAt"}

	var rows [][]string
	for _, row := range data {
		rows = append(rows, []string{
			row["Lot"],
			row["Plate"],
			row["Make"],
			row["Color"],
			row["Size"],
			row["IsHandicap"],
			row["ParkedAt"],
		})
	}

	return utils.ExportToCSV(headers, rows, filename)
}

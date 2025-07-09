package main

import (
	"fmt"
	"os"
	"time"

	"parking-lot/pkg/attendant"
	"parking-lot/pkg/lot"
	"parking-lot/pkg/services"
	"parking-lot/pkg/stratergy"
	"parking-lot/pkg/vehicle"
)

func main() {
	fmt.Println("🚗 Parking Lot System Demo")
	fmt.Println("=========================")

	_ = os.MkdirAll("out", 0755)

	// 1. Create Vehicles
	cars := []*vehicle.Vehicle{
		vehicle.New("DL-1111", "White", "BMW", false, "Small"),
		vehicle.New("DL-2222", "Black", "Audi", false, "Small"),
		vehicle.New("DL-3333", "White", "WagonR", true, "Small"),  // handicap
		vehicle.New("DL-4444", "Red", "Fortuner", false, "Large"), // large
		vehicle.New("DL-5555", "Blue", "Swift", false, "Small"),
	}

	// 2. Create Parking Lots
	lotA := lot.NewParkingLot("Lot A", 2, make(map[string]lot.ParkedVehicle))
	lotB := lot.NewParkingLot("Lot B", 2, make(map[string]lot.ParkedVehicle))
	lots := []*lot.ParkingLot{lotA, lotB}

	// 3. Create Attendant
	att := attendant.NewAttendant("Ravi", lots)

	// 4. Parking Strategies
	evenStrat := &stratergy.ParkEvenly{}
	handicapStrat := &stratergy.HandicapPark{}
	largeStrat := &stratergy.BigVehiclePark{}

	// 5. Park Vehicles
	fmt.Println("\nParking vehicles with different strategies:")
	_, _ = att.ParkWithStrategy(*cars[0], evenStrat)
	_, _ = att.ParkWithStrategy(*cars[1], evenStrat)
	_, _ = att.ParkWithStrategy(*cars[2], handicapStrat)
	_, _ = att.ParkWithStrategy(*cars[3], largeStrat)
	_, _ = att.ParkWithStrategy(*cars[4], evenStrat)

	// 6. Show Lot Status
	fmt.Println("\nLot Status:")
	for _, l := range lots {
		fmt.Printf("%s: %d/%d occupied\n", l.Name, len(l.Vehicles), l.Capacity)
	}

	// 7. Unpark a Vehicle
	fmt.Println("\nUnparking DL-1111 from Lot A:")
	unparked, err := lotA.Unpark("DL-1111")
	if err == nil {
		fmt.Printf("Unparked: %s (%s)\n", unparked.Plate, unparked.Make)
	} else {
		fmt.Println("Unpark error:", err)
	}

	// 8. Pricing Service Demo
	fmt.Println("\nPricing Service Demo:")
	pricing := services.NewPricingService()
	for _, l := range lots {
		for _, pv := range l.GetParkedVehicles() {
			pv.ParkedAt = time.Now().Add(-1 * time.Hour)
			cost := pricing.CalculateCost(pv.Vehicle, pv.ParkedAt)
			fmt.Printf("%s (%s): ₹%.2f\n", pv.Vehicle.Plate, pv.Vehicle.Size, cost)
		}
	}

	// 9. Police Service Demo
	fmt.Println("\nPolice Service Demo:")
	police := services.NewPoliceService(att)
	whiteCars, _ := police.FindCarByColor("White")
	fmt.Println("White cars found:")
	for _, v := range whiteCars {
		fmt.Printf("  %s (%s)\n", v.Plate, v.Make)
	}
	plates := police.GetVehiclePlatesByColor("White")
	fmt.Println("White car plates:", plates)
	lotName, found := police.GetLotByVehicleNumber("DL-3333")
	if found {
		fmt.Println("DL-3333 is parked in:", lotName)
	}
	after := time.Now().Add(-1 * time.Minute)
	recent := police.GetVehiclesParkedAfter(after)
	fmt.Printf("Vehicles parked after %v:\n", after)
	for _, v := range recent {
		fmt.Printf("  %s (%s)\n", v.Plate, v.Make)
	}
	fmt.Println("All parked vehicles summary:")
	for _, entry := range police.GetAllParkedVehiclesSummary() {
		fmt.Printf("  %s: %s (%s %s) in %s\n", entry["Plate"], entry["Make"], entry["Color"], entry["Size"], entry["Lot"])
	}

	// 10. CSV Export Demo
	fmt.Println("\nExporting police queries to CSV (see 'out/' folder):")
	_ = police.ExportPlatesByColorToCSV("White", "out/white_plates.csv")
	_ = police.ExportVehiclesByColorToCSV("White", "out/white_vehicles.csv")
	_ = police.ExportVehiclesParkedAfterToCSV(after, "out/recent_vehicles.csv")
	_ = police.ExportAllParkedVehiclesSummaryToCSV("out/all_vehicles_summary.csv")

	fmt.Println("\n🎉 Demo complete! Check the 'out/' folder for CSV exports.")
}

// Package purchase provides functions related to vehicle purchase decisions and calculations.
package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	bestOption := option1
	if option1 > option2 {
		bestOption = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", bestOption)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	desc := 0.5
	if age < 3 {
		desc = 0.8
	} else if age >= 3 && age < 10 {
		desc = 0.7
	}
	return float64(originalPrice) * desc
}

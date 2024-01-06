package purchase

import "strings"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	ans := strings.Compare(option1, option2)
	if ans == 1 {
		return option2 + " is clearly the better choice."
	}
	return option1 + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if int(age) >= 3 && int(age) < 9 {
		return originalPrice * 0.7
	} else if int(age) < 3 {
		return originalPrice * 0.8
	} else if int(age) >= 10 {
		return originalPrice * 0.5
	}
	return originalPrice
}

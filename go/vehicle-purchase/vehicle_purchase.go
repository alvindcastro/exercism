package purchase

import "strings"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if (kind == "truck") || (kind == "car") {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	output := strings.Compare(option1, option2)
	var returnValue string

	if output < 0 {
		returnValue = option1 + " is clearly the better choice."
	} else if output == 0 {
		returnValue = option1 + " is clearly the better choice."
	} else {
		returnValue = option2 + " is clearly the better choice."
	}

	return returnValue
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var calculatedPrice float64
	if age < 3 {
		calculatedPrice = originalPrice * .80
	} else if age >= 10 {
		calculatedPrice = originalPrice * .50
	} else if age >= 3 && age < 10 {
		calculatedPrice = originalPrice * .70
	}

	return calculatedPrice
}

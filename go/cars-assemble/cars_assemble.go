package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var sP = successRate / 100
	return float64(productionRate) * sP

}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var calculatedWorkingHoursPerHour = CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(calculatedWorkingHoursPerHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	grpsOf10 := carsCount / 10
	remaingCars := carsCount % 10

	return uint((grpsOf10 * 95000) + (remaingCars * 10000))
}

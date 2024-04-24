// Package cars provides functionality related to car production and cost calculation.
package cars

const (
	// ProductionCostForTens represents the cost of producing ten cars together.
	ProductionCostForTens = 95000
	// ProductionCostForSingles represents the cost of producing a single car individually.
	ProductionCostForSingles = 10000
)

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate) * successRate / 100 / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	tens := int(carsCount / 10)
	singles := carsCount % 10
	return uint((tens * ProductionCostForTens) + (singles * ProductionCostForSingles))
}

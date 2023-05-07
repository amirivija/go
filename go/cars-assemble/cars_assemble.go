package cars

import "fmt"

const (
	tensBatchCost int = 95000
	singleCarCost int = 10000
)

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	if productionRate < 0 || successRate > 100 {
		fmt.Println("productionRate ", productionRate, "successRate ", successRate)
		panic("Invalid Inputs.")
	}

	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	if productionRate < 0 || successRate > 100 {
		panic("Invalid Inputs")
	}

	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)

}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {

	if carsCount < 0 {
		panic("carsCount must be greater than zero")
	}
	batchesOfTen := carsCount / 10
	units := carsCount % 10

	return uint(batchesOfTen*tensBatchCost + units*singleCarCost)
}

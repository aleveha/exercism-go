package cars

import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate) / 60 * (successRate / 100))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groups := carsCount / 10
	individuals := carsCount % 10
	return uint(groups*95000 + individuals*10000)
}

func Run() {
	fmt.Println("CalculateWorkingCarsPerHour(1547, 90): ", CalculateWorkingCarsPerHour(1547, 90))
	fmt.Println("CalculateWorkingCarsPerMinute(1105, 90): ", CalculateWorkingCarsPerMinute(1105, 90))
	fmt.Println("CalculateCost(37): ", CalculateCost(21))
}

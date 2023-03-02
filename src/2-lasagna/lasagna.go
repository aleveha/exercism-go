package lasagna

import "fmt"

const OvenTime = 40

func RemainingOvenTime(actual int) int {
	return OvenTime - actual
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers int, actualMinutesInOver int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOver
}

func Run() {
	var actualMinutesInOver = 10
	numberOfLayers := 0
	fmt.Println("Remaining oven time:", RemainingOvenTime(actualMinutesInOver))
	fmt.Println("Preparation time:", PreparationTime(numberOfLayers))
	fmt.Println("Elapsed time:", ElapsedTime(numberOfLayers, actualMinutesInOver))
}

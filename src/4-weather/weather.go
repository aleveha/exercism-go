// Package weather provides tool to get info about specified weather in specified location.
package weather

import "fmt"

// CurrentCondition represents a current weather condition.
var CurrentCondition string

// CurrentLocation represents a current location.
var CurrentLocation string

// Forecast returns a string value that represents the specified weather condition at the specified location. Both values are passed as parameters.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

func Run() {
	fmt.Println(Forecast("Liberec", "Sunny"))
}

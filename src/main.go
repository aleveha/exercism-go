package main

import (
	helloWorld "exercism-go/src/1-hello-world"
	lasagna "exercism-go/src/2-lasagna"
	escape "exercism-go/src/3-escape"
	weather "exercism-go/src/4-weather"
	cars "exercism-go/src/5-cars"
	"fmt"
)

func main() {
	// 1. Hello world
	fmt.Println("1. Hello world")
	helloWorld.Run()
	fmt.Println()

	// 2. Lasagna
	fmt.Println("2. Lasagna")
	lasagna.Run()
	fmt.Println()

	// 3. Escape
	fmt.Println("3. Escape")
	escape.Run()
	fmt.Println()

	// 4. Weather
	fmt.Println("4. Weather")
	weather.Run()
	fmt.Println()

	// 5. Cars
	fmt.Println("5. Cars")
	cars.Run()
	fmt.Println()
}

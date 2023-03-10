package main

import (
	helloWorld "exercism-go/src/1-hello-world"
	speed "exercism-go/src/10-speed"
	cards "exercism-go/src/11-cards"
	lasagna "exercism-go/src/2-lasagna"
	escape "exercism-go/src/3-escape"
	weather "exercism-go/src/4-weather"
	cars "exercism-go/src/5-cars"
	techpalace "exercism-go/src/6-techpalace"
	partyrobot "exercism-go/src/7-partyrobot"
	purchase "exercism-go/src/8-purchase"
	blackjack "exercism-go/src/9-blackjack"
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

	// 6. Welcome
	fmt.Println("6. Welcome")
	techpalace.Run()
	fmt.Println()

	// 7. Party robot
	fmt.Println("7. Party robot")
	partyrobot.Run()
	fmt.Println()

	// 8. Vehicle Purchase
	fmt.Println("8. Vehicle Purchase")
	purchase.Run()
	fmt.Println()

	// 9. Blackjack
	fmt.Println("9. Blackjack")
	blackjack.Run()
	fmt.Println()

	// 10. Speed
	fmt.Println("10. Speed")
	speed.Run()
	fmt.Println()

	// 11. Cards
	fmt.Println("11. Cards")
	cards.Run()
	fmt.Println()
}

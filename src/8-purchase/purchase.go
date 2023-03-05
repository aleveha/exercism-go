package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var result string
	if option1 < option2 {
		result = option1
	} else {
		result = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", result)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * 0.8
	}
	if age >= 10 {
		return originalPrice * 0.5
	}
	return originalPrice * 0.7
}

func Run() {
	fmt.Println("NeedsLicense(\"car\"): ", NeedsLicense("car"))
	fmt.Println("NeedsLicense(\"truck\"): ", NeedsLicense("truck"))
	fmt.Println("NeedsLicense(\"bike\"): ", NeedsLicense("bike"))

	fmt.Println("ChooseVehicle(\"Toyota\", \"Honda\"): ", ChooseVehicle("Toyota", "Honda"))

	fmt.Println("CalculateResellPrice(1000, 1): ", CalculateResellPrice(1000, 1))
	fmt.Println("CalculateResellPrice(1000, 5): ", CalculateResellPrice(1000, 5))
	fmt.Println("CalculateResellPrice(1000, 15): ", CalculateResellPrice(1000, 15))
}

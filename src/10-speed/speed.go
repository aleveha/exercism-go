package speed

import "fmt"

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
	}
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery < car.batteryDrain {
		return car
	}

	car.distance += car.speed
	car.battery -= car.batteryDrain
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	movesCount := track.distance / car.speed
	if track.distance%car.speed != 0 {
		movesCount += 1
	}
	return car.battery >= car.batteryDrain*movesCount
}

func Run() {
	car := NewCar(10, 2)
	track := NewTrack(100)
	fmt.Printf("NewCar(10,2): %v\n", car)
	fmt.Printf("NewTrack(100): %v\n", track)
	fmt.Printf("Drive(car): %v\n", Drive(car))
	fmt.Println("CanFinish(Car{battery: 100, batteryDrain: 5, speed: 5, distance: 20}, Track{distance: 100}): ", CanFinish(Car{battery: 100, batteryDrain: 5, speed: 5, distance: 20}, Track{distance: 100}))
	fmt.Println("CanFinish(Car{battery:25, batteryDrain:3, speed:2, distance:0}, Track{distance: 16}): ", CanFinish(Car{battery: 25, batteryDrain: 3, speed: 2, distance: 0}, Track{distance: 16}))
}

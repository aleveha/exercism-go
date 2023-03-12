package birdwatcher

import "fmt"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var counter int
	for i := 0; i < len(birdsPerDay); i++ {
		counter += birdsPerDay[i]
	}
	return counter
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var counter int
	for i := (week - 1) * 7; i < week*7; i++ {
		counter += birdsPerDay[i]
	}
	return counter
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}

func Run() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}

	fmt.Println("TotalBirdCount: ", TotalBirdCount(birdsPerDay))
	fmt.Println("BirdsInWeek: ", BirdsInWeek(birdsPerDay, 2))
	fmt.Println("FixBirdCountLog: ", FixBirdCountLog(birdsPerDay[:6]))
}

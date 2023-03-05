package partyrobot

import (
	"fmt"
	"strconv"
	"strings"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var tableNumberAsString string
	if table < 100 {
		var zerosCount int
		if table < 10 {
			zerosCount = 2
		} else {
			zerosCount = 1
		}
		tableNumberAsString = fmt.Sprintf("%s%d", strings.Repeat("0", zerosCount), table)
	} else {
		tableNumberAsString = strconv.Itoa(table)
	}

	welcome := fmt.Sprintf("Welcome to my party, %s!", name)
	tableInfo := fmt.Sprintf("You have been assigned to table %s. Your table is %s, exactly %.1f meters from here.", tableNumberAsString, direction, distance)
	seatmate := fmt.Sprintf("You will be sitting next to %s.", neighbor)
	return strings.Join([]string{welcome, tableInfo, seatmate}, "\n")
}

func Run() {
	name := "aleveha"
	age := 22

	fmt.Println(Welcome(name))
	fmt.Println(HappyBirthday(name, age))
	fmt.Println(AssignTable(name, 10, "Mike", "on the left", 23.448362))
}

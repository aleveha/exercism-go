package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := strings.Repeat("*", numStarsPerLine)
	return fmt.Sprintf("%s\n%s\n%s", stars, welcomeMsg, stars)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	withoutStars := strings.Replace(oldMsg, "*", "", -1)
	return strings.TrimSpace(withoutStars)
}

func Run() {
	welcomeMessage := WelcomeMessage("aleveha")
	fmt.Println(welcomeMessage)
	fmt.Println(AddBorder(welcomeMessage, 20))
	fmt.Println(CleanupMessage(AddBorder(welcomeMessage, 20)))
}

package blackjack

import "fmt"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)

	if card1Value == 11 && card2Value == 11 {
		return "P"
	}

	cardsSum := card1Value + card2Value

	if cardsSum == 21 {
		if dealerCardValue < 10 {
			return "W"
		} else {
			return "S"
		}
	}

	if cardsSum >= 17 && cardsSum <= 21 {
		return "S"
	}

	if cardsSum >= 12 && cardsSum <= 16 {
		if dealerCardValue >= 7 {
			return "H"
		} else {
			return "S"
		}
	}

	return "H"
}

func Run() {
	fmt.Println("ParseCard(ace): ", ParseCard("ace"))
	fmt.Println("ParseCard(six): ", ParseCard("six"))
	fmt.Println("FirstTurn(ace,ace,ace): ", FirstTurn("ace", "ace", "ace"))
	fmt.Println("FirstTurn(ace,ten,nine): ", FirstTurn("ace", "ten", "nine"))
	fmt.Println("FirstTurn(ace,ten,ten): ", FirstTurn("ace", "ten", "ten"))
	fmt.Println("FirstTurn(ten,nine,ten): ", FirstTurn("ten", "nine", "nine"))
	fmt.Println("FirstTurn(five,eight,nine): ", FirstTurn("five", "eight", "nine"))
	fmt.Println("FirstTurn(five,eight,two): ", FirstTurn("five", "eight", "two"))
	fmt.Println("FirstTurn(two,four,ten): ", FirstTurn("two", "four", "nine"))
}

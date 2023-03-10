package cards

import "fmt"

func isIndexInvalid(sliceLength, index int) bool {
	return index < 0 || index >= sliceLength
}

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if isIndexInvalid(len(slice), index) {
		return -1
	}
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if isIndexInvalid(len(slice), index) {
		return append(slice, value)
	}

	slice[index] = value
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	if len(values) == 0 {
		return slice
	}
	return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if isIndexInvalid(len(slice), index) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func Run() {
	cards := FavoriteCards()
	fmt.Println(cards)
	println("1–––––")

	getCardSuccess := GetItem(cards, 1)
	fmt.Println(getCardSuccess)
	getCardError := GetItem(cards, -12)
	fmt.Println(getCardError)
	println("2–––––")

	setCardSuccess := SetItem(cards, 1, 3)
	fmt.Println(setCardSuccess)
	setCardError := SetItem(cards, 10, 3)
	fmt.Println(setCardError)
	println("3–––––")

	prependItemsWithValues := PrependItems(cards, 5, 1, 6)
	fmt.Println(prependItemsWithValues)
	prependItemsWithoutValues := PrependItems(cards)
	fmt.Println(prependItemsWithoutValues)
	println("4–––––")

	removeItemsSuccess := RemoveItem(cards, 1)
	fmt.Println("removeItemsSuccess: ", removeItemsSuccess)
	fmt.Println("cards: ", cards)
	removeItemsError := RemoveItem(cards, 10)
	fmt.Println(removeItemsError)
	println("5–––––")
}

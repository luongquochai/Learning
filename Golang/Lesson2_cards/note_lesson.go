/*
THIS FILE IS FOR KEEPING EVERY LESSON'S NOTES
*/

package main

import "fmt"

func main() {
	/*
		// var card string = "Ace of Spades"
		card := newCard()
		// := helps define kind of type of variable
		// card = "Five of Diamonds"
	*/
	// Create a slice type string with two elements
	// cards := []string{"Ace of Diamonds", newCard()}
	// Replace []string -> deck
	// cards := deck{"Ace of Diamonds", newCard()}
	// // Add an new element to slice
	// cards = append(cards, "Six of Spades")

	// =====================================================
	// cards := newDeck()

	// // cards.print()

	// // assign a tuple value
	// hand, remainingCards := deal(cards, 7)

	// hand.print()
	// fmt.Println("----------------------------")
	// remainingCards.print()
	// =====================================================
	cards := newDeck()
	fmt.Println(cards.toString())
}

// func newCard() string {
// 	return "Ace of Spades"
// }

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
	cards := []string{"Ace of Diamonds", newCard()}
	// Add an new element to slice
	cards = append(cards, "Six of Spades")

	// Use iterate to print each of elements
	for i, card := range cards {
		fmt.Println(i, ": ", card)
	}
}

func newCard() string {
	return "Ace of Spades"
}

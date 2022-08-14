package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Create a newDeck function
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// ------------------------------------------------------------------
// Base on the intent, we need to define is the function or receiver
// ------------------------------------------------------------------
// this function adds a new method (print) into variable 'd' type 'deck'

// receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// function helper
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// create a new receiver:
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// function helper save to file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// function get newDeck from file
func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)

	if err != nil {
		// Option #1 - Log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(byteSlice), ",")
	return deck(s)
}

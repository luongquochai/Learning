package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

/*
NOTE: 	In order to use the function 'WriteFile' of ioutil package, we need convert 'slice of strings'
		to 'slice of bytes'.
I.E:	slice_of_strings := []string{"A","B","C"}
		-> fmt.Println([]bye(slice_of_strings))
		output: 65 66 67
*/
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

// Create a new Shuffle function, which is helps to change position of value in slice
// To do that, we will iterate length of slice and swap index with a new randomized position
func (d deck) shuffle() {

	// Since 'rand.Intn' will random with only one seed which need to be changed in the loop
	// Therefore, we will use another function can use 'seed' to initialize a new random in loop

	// seed int64 and func (t Time) UnixNano() int64
	source := rand.NewSource(time.Now().UnixNano())
	// func New(src Source) *Rand
	// New returns a new Rand that uses random values from src to generate other random values.
	r := rand.New(source)

	for i := range d {
		// Due to r will return type *Rand have Intn
		newPos := r.Intn(len(d) - 1)

		d[i], d[newPos] = d[newPos], d[i]
	}
}

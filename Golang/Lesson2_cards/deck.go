package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// this function adds a new method (print) into variable 'd' type 'deck'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

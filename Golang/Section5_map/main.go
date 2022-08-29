package main

import (
	"fmt"
)

func main() {

	/*two initial ways of declaring a map

	1. var colors map[string]string

	2. colors := make(map[string]string)

	*/

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b765",
	}

	//Add a pair key and value in exist map
	colors["white"] = "#ffffff"

	// Delete key-value
	// delete(colors, "red")

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code of", key, "is", value)
	}
}

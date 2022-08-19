/*
Here are the directions:

1. Create a new project folder to house this new project and create a 'main.go' file inside of it.
2. In the main.go file, define a function called 'main'.  Remember that the 'main' function will be called automatically.
3. In the main function, create a slice of ints from 0 through 10
4. Iterate through the slice with a for loop.  For each element, check to see if the number is even or odd.
5. If the value is even, print out "even".  If it is odd, print out "odd"
6. Run your code from the terminal by changing into your project directory then running 'go run main.go'.
*/

package main

import "fmt"

func main() {
	sliceOfints := []int{}

	for _, number := range sliceOfints {
		if number%2 == 0 {
			fmt.Println(number, " is even")
		} else {
			fmt.Println(number, " is odd")
		}
	}
}

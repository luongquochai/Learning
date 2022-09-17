package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

// method
func (englishBot) getGreeting() string {
	return "Hi Am English Bot"
}

func (spanishBot) getGreeting() string {
	return "Hi Am Spanish Bot"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

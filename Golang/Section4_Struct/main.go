package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// First way:
	// Hai := person{"Hai","Luong"} need to be ordered in struct
	// =========================================
	// Second way:
	// Hai := person{firstName: "Hai",
	// 	lastName: "Luong",
	// }
	// fmt.Println(Hai)
	// =========================================
	// Third way:
	var Hai person
	Hai.firstName = "Hai"
	Hai.lastName = "Luong"
	fmt.Println(Hai)
	fmt.Printf("%+v", Hai)
}

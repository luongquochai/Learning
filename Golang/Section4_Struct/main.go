package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
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
	// var Hai person
	// Hai.firstName = "Hai"
	// Hai.lastName = "Luong"
	// fmt.Println(Hai)
	// fmt.Printf("%+v", Hai)

	/* Embedded Struct */
	Hai := person{
		firstName: "Hai",
		lastName:  "Luong",
		contact: contactInfo{
			email:   "hai@gmail.com",
			zipCode: 70000,
		},
	}

	// HaiPointer := &Hai
	// HaiPointer.updateName("Haiii")
	/*
		HaiPointer = type of * person, or a pointer to a person
		Hai = type of person
	*/
	Hai.updateName("Haiii")
	Hai.print()

}

/*
NOTE:
*person: this is a type description, it means we're working with a pointer to a person
*pointerToPerson: this is an operator, it means we want to manipute the value the pointer is referencing
---- shortcut pointer
*person: type of * person, or a pointer to a person
*/

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

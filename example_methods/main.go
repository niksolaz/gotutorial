package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Surname string
}

func (n *Person) UpdateName(newN string) {
	n.Name = newN
}

func (n *Person) UpdateSurname(newN string) {
	n.Surname = newN
}

func (n *Person) PrintFullName() {
	fmt.Println(n)
}

func myFullName(firstname string, lastname string) (string, error) {
	fullname := firstname + " " + lastname
	return fullname, nil
}

func main() {
	var person Person
	person.Name = "Nicola"
	person.Surname = "Solazzo"
	f, err := myFullName(person.Name, person.Surname)

	if err != nil {
		fmt.Println("Handle error")
	}

	fmt.Println("----------------")
	fmt.Println("Actual Fullname: ", f)
	fmt.Println("----------------")

	person.UpdateName("Giulio")
	person.UpdateSurname("Cesare")
	person.PrintFullName()
	fmt.Println("----------------")
	fmt.Println("Update Fullname: ", person.Name, person.Surname)
	fmt.Println("----------------")

}

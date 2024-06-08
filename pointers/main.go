package main

import (
	"fmt"
)

// Engeneer stores a name and age
type Engeneer struct {
	Name string
	Age  int
}

func (e *Engeneer) Birthday() {
	e.Age++
	fmt.Println("Birthday!", e)
}

func (e Engeneer) UpdateName(name string) {
	e.Name = name
	fmt.Println(e)
}

func UpdateName(e *Engeneer, name string) {
	e.Name = name
	fmt.Println(e)
}

func main() {
	fmt.Println("pointers")

	// var name string = "John"
	// fmt.Println(name)

	// ptr := &name
	// fmt.Println(*ptr)

	// *ptr = "Jane"
	// fmt.Println(name)

	lolo := &Engeneer{
		Name: "Lolo",
		Age:  25,
	}
	fmt.Println(lolo)
	lolo.Birthday()
	fmt.Println(lolo)
	lolo.UpdateName("Lola")
	fmt.Println(lolo)
	UpdateName(lolo, "Lola")
	fmt.Println("Lola", lolo)
}

package main

import (
	"errors"
	"fmt"
)

func DoSomething() (err error) {
	defer func() {
		err = errors.New("error")
	}()
	fmt.Println("doSomething")
	return nil
}

func ReturnValues() (v int) {
	defer func() {
		v = 1
	}()
	v = 5

	return
}

type Engeneer struct {
	Name string
}

func (e *Engeneer) ChangeName(name string) {
	e.Name = name
}

func doStuff(e *Engeneer) {
	name := "John Doe"
	defer e.ChangeName(name)
	fmt.Println("doStuff")
	name = "John Doe Doe"
}

func main() {

	fmt.Println("Defer")

	john := &Engeneer{
		Name: "John",
	}

	fmt.Printf("%+v\n", john)
	doStuff(john)
	fmt.Printf("%+v\n", john)

	// defer func() {
	// 	fmt.Println("main")
	// }()

	// err := doSomething()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(ReturnValues())
}

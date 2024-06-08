package main

import (
	"errors"
	"fmt"
)

func IPanic() {
	panic("aaaaaaaaaaaaa")
}

func doSomething() (err error) {
	fmt.Println("doSomething")

	defer func() {
		fmt.Println("2")
		if r := recover(); r != nil {
			fmt.Println("Recovered!")
			err = errors.New("IPanic panicked")
		}
	}()
	IPanic()
	return nil
}

func main() {
	fmt.Println("Panic!")
	if err := doSomething(); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("main")
}

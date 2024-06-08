package main

import (
	"errors"
	"fmt"
)

var (
	ErrIsLessThanZero = errors.New("number is less than zero")
	ErrIsNotEven      = errors.New("number is not even")
)

func IsEven(n int) error {
	if n%2 != 0 {
		return ErrIsNotEven
	}
	return nil
}

func main() {
	fmt.Println("Errors")

	// IsEven(1) bad practice
	IsEven(1)
	err := IsEven(22)
	if err != nil {
		fmt.Println("number is not even")
	}

	err = IsEven(1)
	if err != nil {
		if err == ErrIsNotEven {
			fmt.Println("err is not even throw")
		}
		fmt.Println("111 number is not even")
		fmt.Println(err.Error())
	}
}

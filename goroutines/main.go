package main

import (
	"fmt"
	"math/rand"
)

// func HelloWorld(name string) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Hello", name)
// }

func CalculateValues(values chan int) {
	random := rand.Intn(100)
	fmt.Println("random is", random)
	values <- random
}

func main() {

	values := make(chan int)
	// go HelloWorld("Lolo")
	go CalculateValues(values)

	value := <-values
	fmt.Println("value is", value)

	// fmt.Println("I should be printed after 2 seconds")
	// time.Sleep(2 * time.Second)

}

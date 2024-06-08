package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValues(v chan int) {
	for i := 0; i <= 10; i++ {
		random := rand.Intn(100)
		fmt.Println("random is", random)
		v <- random
	}
}

func main() {

	values := make(chan int, 2)
	go CalculateValues(values)

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		value := <-values
		fmt.Println("value is", value)
	}
}

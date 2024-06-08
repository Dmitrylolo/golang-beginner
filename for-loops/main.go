package main

import "fmt"

func main() {
	fmt.Println("For Loops")

	ages := []int{22, 36, 77, 88}

	for index, value := range ages {
		fmt.Println(index, value)
	}
	fmt.Println("====================")

	for i := 0; i < len(ages); i++ {
		fmt.Println(ages[i])
	}
	fmt.Println("====================")

	for i := 0; i < len(ages); {
		fmt.Println(ages[i])
		i++
	}

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

}

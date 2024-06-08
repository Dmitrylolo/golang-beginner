package main

import "fmt"

func HelloWorld(name string, age, height int) {
	fmt.Println("Hello, " + name + "!")
	fmt.Println("You are", age)
	fmt.Println("Your height is", height)
}

// func AddTotal(a, b int) int {
// 	return a + b
// }

func AddTotal(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	fmt.Println("Functions")
	name := "Kuku"
	var age int = 22
	height := 178
	HelloWorld(name, age, height)
	totalPlus, totalMinus := AddTotal(age, height)
	// fmt.Println("Total is", total)
	fmt.Println("Total is", totalPlus, totalMinus)
}

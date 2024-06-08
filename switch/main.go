package main

import "runtime"

func main() {
	println("If statement")

	var customerHeight int = 178
	customerAge := 18

	switch {
	case customerHeight >= 150 && customerAge >= 18:
		println("Can Access ride")
	case customerHeight >= 120:
		println("Customer can access children's ride")
	default:
		println("Customer is too small")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		println("Mac OS")
	case "linux":
		println("Linux")
	case "windows":
		println("Windows")
	default:
		println("Unknown OS")
	}

	// if customerHeight >= 150 || customerAge >= 18 {
	// 	println("Can Access ride")
	// } else if customerHeight >= 120 {
	// 	println("Customer can access children's ride")
	// } else {
	// 	println("Customer is too small")
	// }
}

package main

func main() {
	println("If statement")

	var customerHeight int = 178
	customerAge := 11

	if customerHeight <= 180 && customerAge >= 18 {
		println("Can Access full service")
	} else if customerHeight <= 210 {
		println("Customer cann't access full service")
	} else {
		println("Customer is over 210 cm")
	}
}

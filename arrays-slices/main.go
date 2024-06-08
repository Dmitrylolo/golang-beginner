package main

import "fmt"

func main() {
	fmt.Println("Arrays and Slices")

	plantes := [8]string{"mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "neptune"}
	fmt.Println(plantes)

	var planetsArray [8]string
	planetsArray[0] = "mercury"
	fmt.Println(planetsArray)

	planetsSlice := []string{"mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "neptune"}
	fmt.Println(planetsSlice)

	var planetsSliceVerbose []string
	planetsSliceVerbose = append(planetsSliceVerbose, "mercury")
	planetsSliceVerbose = append(planetsSliceVerbose, "venus")
	planetsSliceVerbose = append(planetsSliceVerbose, "earth")
	fmt.Println(planetsSliceVerbose)
}

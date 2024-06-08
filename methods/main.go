package main

import "fmt"

type Engeneer struct {
	Name    string
	Age     int
	Project Project
}

type Project struct {
	Name       string
	Priority   string
	Technology []string
}

func (e Engeneer) Print() {
	fmt.Println("Engeneer info")
	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Age: %d\n", e.Age)
	fmt.Printf("Current Project: %s\n", e.Project.Name)
}

func (e *Engeneer) BirthDay() {
	e.Age++
	fmt.Printf("Birthday!")
}

func (e *Engeneer) GetProjectName() string {
	return e.Project.Name
}

func main() {
	fmt.Println("Methods")
	engeneer := Engeneer{
		Name: "lolo",
		Age:  20,
		Project: Project{
			Name:       "go",
			Priority:   "medium",
			Technology: []string{"go"},
		},
	}

	// project := Project{
	// 	Name:       "gov",
	// 	Priority:   "hifi",
	// 	Technology: []string{"go1", "go2"},
	// }
	engeneer.Print()
	engeneer.BirthDay()
	engeneer.Print()

	fmt.Println(engeneer.GetProjectName())

	// project.Print()
}

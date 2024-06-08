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

func main() {
	fmt.Println("Structs")
	engeneer := Engeneer{
		Name: "lolo",
		Age:  20,
		Project: Project{
			Name:       "go",
			Priority:   "medium",
			Technology: []string{"go"},
		},
	}
	fmt.Printf("%+v\n", engeneer)

	fmt.Println(engeneer.Name)
	fmt.Println(engeneer.Project.Priority)
}

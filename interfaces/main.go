package main

import "fmt"

type Employee interface {
	GetName() string
}

type Engeneer struct {
	Name string
}

func (e *Engeneer) GetName() string {
	return "Name: " + e.Name
}

type Manager struct {
	Name string
}

func (m *Manager) GetName() string {
	return "Name: " + m.Name
}

func PrintDetails(e Employee) {
	fmt.Println(e.GetName())
}

func main() {
	fmt.Println("Interfaces")
	engeneer := &Engeneer{
		Name: "Lolo",
	}

	manager := &Manager{
		Name: "TT",
	}

	PrintDetails(engeneer)
	PrintDetails(manager)

}

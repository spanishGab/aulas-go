package basics

import (
	"fmt"
	"time"
)

// Uma struct é um tipo de dado composto que agrega multiplos valores de tipos diferentes
type Animal struct {
	Breed       string
	Name        string
	Age         int
	IsCastrated bool
	Vaccines    map[string]time.Time
}

type Point struct {
	X int
	Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func Structs() {
	structBasics()
	structComparison()
	structComposition()
}

func structBasics() {
	animal := Animal{
		Breed:       "Shihtzu",
		Name:        "Thor",
		Age:         14,
		IsCastrated: true,
		Vaccines: map[string]time.Time{
			"rage": time.Now(),
		},
	}
	// Por convenção, structs muito grandes são passadas como ponteiro para funções

	fmt.Printf("animal=%#v\n", String(&animal))

	thor := &animal

	animal.Breed = "No Breed"
	animal.Name = "Safira"
	fmt.Printf("thor.Name=%s, animal.Name=%s\n", thor.Name, animal.Name)
}

func structComparison() {
	// É possível comparar strcuts cujos tipos são todos comparáveis
	a := Point{1, 2}
	b := Point{1, 2}
	fmt.Printf("a==b=%t\n", a == b)
}

func structComposition() {
	circle := Circle{Point{5, 5}, 3}
	wheel := Wheel{circle, 6}
	fmt.Printf("wheel=%#v\n", wheel)
}

func String(animal *Animal) string {
	return fmt.Sprintf("Breed: %s, Name: %s, Age: %d, IsCastrated: %t, Vacccines: %#v\n",
		animal.Breed,
		animal.Name,
		animal.Age,
		animal.IsCastrated,
		animal.Vaccines,
	)
}

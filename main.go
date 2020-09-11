package main

import "fmt"

// Animal struct
type Animal struct {
	Fly  bool
	Legs int
	Name string
}

// NewAnimal is the Animal struct constructor
func NewAnimal(fly bool, legs int, name string) *Animal {
	return &Animal{
		Fly:  fly,
		Legs: legs,
		Name: name,
	}
}

func (a *Animal) tryFly() string {
	if a.Fly {
		return fmt.Sprintf("It flies")
	}
	return fmt.Sprintf("A %s can't fly!", a.Name)
}

func main() {
	fmt.Println("Hello World")

	var bird *Animal
	bird = NewAnimal(true, 2, "eagle")
	fmt.Printf("%s has %d legs and %s\n", bird.Name, bird.Legs, bird.tryFly())

	var duck *Animal
	duck = NewAnimal(false, 2, "donald")
	fmt.Printf("a %s Can fly? %s\n", duck.Name, duck.tryFly())
}

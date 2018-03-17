package main

import (
	"fmt"
)

func main() {

	fruit := Fruit{
		"Apple",
		"Red",
	}

	fmt.Printf("Fruit Memory Address Before: %p\n", &fruit)

	passByValue(fruit)

	fmt.Println("Fruit After by Value: ", fruit)

	passByReferece(&fruit)

	fmt.Println("Fruit After by Reference: ", fruit)

}

func passByValue(fruit Fruit) Fruit {
	fruit.name = fruit.name + "changed"
	return fruit
}

func passByReferece(fruit *Fruit) *Fruit {
	fruit.name = fruit.name + "changed"
	return fruit
}

type Fruit struct {
	name  string
	color string
}

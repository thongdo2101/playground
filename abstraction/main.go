package main

import "fmt"

type Flyable interface {
	Fly() string
}

type FlyableAnimal struct {
	Flyable
}

func MakeFly(animal Flyable) string {
	return animal.Fly()
}

type Duck struct {
}

func (d Duck) Fly() string {
	return "duck flying"
}

type Butterfly struct {
}

func (b Butterfly) Fly() string {
	return "butterfly flying"
}

func main() {
	duck := Duck{}
	butterfly := Butterfly{}
	duckAnimal := FlyableAnimal{duck}
	butterflyAnimal := FlyableAnimal{butterfly}
	fmt.Println(MakeFly(duckAnimal))
	fmt.Println(MakeFly(butterflyAnimal))
	// MakeFly chỉ biết rằng 1 FlyableAnimal truyền vào mà không biết nó là gì
}

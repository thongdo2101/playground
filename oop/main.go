package main

import "fmt"

func main() {
	b := new(B)
	b.Name = "b"
	b.Print()
}

type A struct {
	Name string
}

func (a A) Print() {
	fmt.Println(a.Name)
}

type B struct {
	A
}

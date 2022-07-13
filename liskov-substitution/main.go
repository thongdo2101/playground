package main

import "fmt"

type Employee struct {
	Name string
}

func (e *Employee) GetName() {
	fmt.Println(e.Name)
}

type Developer struct {
	Employee
}

func (d *Developer) Develop() {
	fmt.Println("developing")
}

func (d *Developer) GetName() {
	fmt.Println(d.Name)
}

type Tester struct {
	Employee
}

func (t *Tester) Test() {
	fmt.Println("testing")
}

func main() {
	d := Developer{Employee{"John"}}
	d.GetName()
	d.Develop()
	t := Tester{Employee{"Jane"}}
	t.GetName()
}

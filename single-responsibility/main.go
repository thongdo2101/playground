package main

import "fmt"

type Calculator struct{}

func (c *Calculator) Calculate(a, b int, operator string) int {
	if operator == "+" {
		return a + b
	}
	if operator == "-" {
		return a - b
	}
	if operator == "*" {
		return a * b
	}
	if operator == "/" {
		return a / b
	}
	return 0
}

func (c *Calculator) Add(a, b int) int {
	return a + b
}
func (c *Calculator) Subtract(a, b int) int {
	return a - b
}
func (c *Calculator) Multiply(a, b int) int {
	return a * b
}
func (c *Calculator) Divide(a, b int) int {
	return a / b
}
func (c *Calculator) Calculate2(a, b int, operator string) int {
	switch operator {
	case "+":
		return c.Add(a, b)
	case "-":
		return c.Subtract(a, b)
	case "*":
		return c.Multiply(a, b)
	case "/":
		return c.Divide(a, b)
	}
	return 0
}

//

type EmployeeNotSingle struct {
	Role string
}

func (d *EmployeeNotSingle) Develop() {
	fmt.Println("developing")
}

func (t *EmployeeNotSingle) Test() {
	fmt.Println("testing")
}

//

type Employee struct {
	Name string
}

type Developer struct {
	Employee
}

func (d *Developer) Develop() {
	fmt.Println("developing")
}

type Tester struct {
	Employee
}

func (t *Tester) Test() {
	fmt.Println("testing")
}

func main() {

}

package main

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

func (c * Calculator) Add(a, b int) int {
  return a + b
}
func (c * Calculator) Subtract(a, b int) int {
  return a - b
}
func (c * Calculator) Multiply(a, b int) int {
  return a * b
}
func (c * Calculator) Divide(a, b int) int {
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

func main() {

}

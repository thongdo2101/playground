package main

import "fmt"

func main() {
	str := "12345"
	newStr := str[3:]
	fmt.Println(str)
	fmt.Println(newStr)
	fmt.Println(&str)
	fmt.Println(&newStr)
}

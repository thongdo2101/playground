package main

import "fmt"

func main() {
	realVar := 1
	array := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	for _, v := range array {
		realVar := v
		fmt.Println(realVar)
	}
	fmt.Println("---")
	fmt.Println(realVar)
}

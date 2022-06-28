package main

import "fmt"

func main() {
	i := make([]int, 0)
	fmt.Println(len(i))
	fmt.Println(i == nil)
	fmt.Println("--")
	// var j *int
	// fmt.Println(j)
	// *j = 5
	// fmt.Println(j)
	var m map[int]int
	fmt.Println("--")
	fmt.Println(m)
	fmt.Println(m == nil)
	fmt.Println("--")
	m2 := make(map[int]int)
	fmt.Println(m2)
	fmt.Println(m2 == nil)
}

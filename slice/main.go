package main

import "fmt"

func main() {
	// slice1 := []int{1, 2}
	// slice2 := []int{3, 4}
	// fmt.Println(slice1, slice2)
	// fmt.Println(fmt.Sprintf("%p %p %p", slice1, slice2))
	// slice3 := slice1
	// slice1 = slice2
	// // copy(slice1, slice2)
	// fmt.Println(slice1, slice2, slice3)
	// fmt.Println(fmt.Sprintf("%p %p %p", slice1, slice2, slice3))
	//
	slice := make([]int, 0)
	fmt.Println(fmt.Sprintf("%p", slice))
	newSlice := append(slice, 1)
	fmt.Println(fmt.Sprintf("%p", newSlice))
	newSlice = append(slice, 1)
	fmt.Println(fmt.Sprintf("%p", newSlice))
}

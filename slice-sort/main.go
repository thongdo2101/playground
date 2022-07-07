package main

import (
	"fmt"
	"sort"
)

// type MySort interface {
// 	Len() int
// 	Less(i, j int) bool
// 	Swap(i, j int)
// }

type Values []int

func (list Values) Len() int {
	return len(list)
}
func (list Values) Less(i, j int) bool {
	return list[i] < list[j]
}
func (list *Values) Swap(i, j int) {
	(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
}

func main() {
	values := Values{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	sort.Sort(&values)
	fmt.Println(values)
}

package main

import (
	"fmt"

	"github.com/volatiletech/null"
)

func main() {
	var str null.String
	fmt.Println(str)
	fmt.Println(str.String)
}

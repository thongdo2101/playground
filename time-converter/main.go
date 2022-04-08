package main

import (
	"fmt"

	"github.com/araddon/dateparse"
)

func main() {
	// layout := "2022-03-02T09:02:53+00:00"
	str := "2022-03-02T09:02:53+00:00"
	t, err := dateparse.ParseLocal(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
}

package main

import (
	"fmt"
	"golangplayground/scope/internal"
)

func main() {
  fmt.Println("start")
	module := new(internal.Module)
	fmt.Println(module)
	module.Name = "Hello, World!"
	module.SetPrivateName("xxx")
	fmt.Println(module)
	module.SetPrivateName2("xxx222")
	fmt.Println(module)
}

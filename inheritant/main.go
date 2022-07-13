package main

import "fmt"

type A struct {
}

type B struct {
	A
}

func test(a A) {
	fmt.Println("A")
}

type IA interface {
	Foo()
}

type B2 struct {
	IA
}

func (b2 B2) Foo() {
	fmt.Println("B2")
}

func test2(i IA) {
	i.Foo()
}

func main() {
	// b := B{}
	// test(b) // cannot use b (type B) as type A in argument to test:
	// B bao gồm A chứ không phải A
	b2 := B2{}
	test2(b2)
}

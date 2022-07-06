package internal

import "fmt"

func init() {
	fmt.Println("xxx")
}

var internalString = "Hello, World!"

type Module struct {
	Name        string
	privateName string
}

func (m *Module) SetPrivateName(name string) {
	m.privateName = name
}
func (m Module) SetPrivateName2(name string) {
	m.privateName = name
}

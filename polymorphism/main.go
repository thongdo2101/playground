package main

type IGun interface {
	Shoot()
}

type AK47 struct {
}

func (ak47 AK47) Shoot() {
	println("AK47")
}

type AWM struct {
}

func (awm AWM) Shoot() {
	println("AWM")
}

func test(gun IGun) {
	gun.Shoot()
}

func main() {
	ak47 := AK47{}
	awm := AWM{}
	test(ak47)
	test(awm)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	c := make(chan bool)
	go func(c chan bool) {
		time.Sleep(time.Second * 2)
		fmt.Println("go")
		c <- true // send value to channel
	}(c)
	<-c
	fmt.Println("end")
}

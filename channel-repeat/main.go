package main

import (
	"fmt"
	"time"
)

func main() {
	receiver := make(chan string, 2)
	go func(respondent chan string) {
		for {
			fmt.Println("response: " + <-respondent)
		}
	}(receiver)
	for {
		input := ""
		fmt.Println("input: ")
		fmt.Scanf("%s", &input)
		receiver <- input
		time.Sleep(time.Second / 3)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	receiver := make(chan string, 2)
	go func(respondent chan string) {
		for {
			time.Sleep(time.Second * 1)
			fmt.Println("response: " + <-respondent)
		}
	}(receiver)
	for {
		fmt.Println("input: ")
		in := bufio.NewReader(os.Stdin)
		line, _ := in.ReadString('\n')
		receiver <- line
		time.Sleep(time.Second / 3)
	}
}

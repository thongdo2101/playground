package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	numberOfURLs    = 100
	numberOfWorkers = 5
)

func crawlURL(queue <-chan int, name string, quit chan<- int) {
	counted := 0
	for v := range queue {
		fmt.Printf("Worker %s is crawling URL %d. total: %d\n", name, v, counted)
		time.Sleep(time.Second)
		counted++
	}
	numb, _ := strconv.ParseInt(name, 10, 10)
	quit <- int(numb)
	fmt.Printf("Worker %s done and exit\n", name)
}

func main() {
	queue := make(chan int, 100)
	quit := make(chan int, numberOfWorkers)
	go func() {
		for i := 1; i <= numberOfURLs; i++ {
			queue <- i
			fmt.Printf("URL %d has been enqueued\n", i)
		}

		close(queue)
	}()

	for i := 1; i <= numberOfWorkers; i++ {
		go crawlURL(queue, fmt.Sprintf("%d", i), quit)
	}
	for i := 1; i <= numberOfWorkers; i++ {
		<-quit
	}
}

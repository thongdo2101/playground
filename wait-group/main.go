package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numberOfURLs    = 100
	numberOfWorkers = 5
)

func crawlURL(queue <-chan int, name string, wg sync.WaitGroup) {
	counted := 0
	for v := range queue {
		fmt.Printf("Worker %s is crawling URL %d. total: %d\n", name, v, counted)
		time.Sleep(time.Second)
		counted++
	}
	wg.Done()
	fmt.Printf("Worker %s done and exit\n", name)
}

func main() {
	queue := make(chan int, 100)
	wg := new(sync.WaitGroup)
	wg.Add(numberOfWorkers + 1)
	go func() {
		for i := 1; i <= numberOfURLs; i++ {
			queue <- i
			fmt.Printf("URL %d has been enqueued\n", i)
		}
		wg.Done()
		close(queue)
	}()

	for i := 1; i <= numberOfWorkers; i++ {
		go crawlURL(queue, fmt.Sprintf("%d", i), *wg)
	}
	wg.Wait()
	fmt.Println("All workers done and exit")
}

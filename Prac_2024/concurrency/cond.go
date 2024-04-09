package main

import (
	"fmt"
	"sync"
)

var sharedData int
var cond = sync.NewCond(&sync.Mutex{})
var producerFinished bool // Indicates whether the producer has finished

func producer(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		cond.L.Lock()
		sharedData = i
		fmt.Println("Produced:", sharedData)
		cond.Signal() // Signal to consumer that new data is available
		cond.L.Unlock()
	}
	// Mark producer as finished
	cond.L.Lock()
	producerFinished = true
	cond.Broadcast() // Signal all waiting goroutines that producer has finished
	cond.L.Unlock()
}

func consumer(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		cond.L.Lock()
		for sharedData == -1 { // Wait until new data is available
			if producerFinished {
				cond.L.Unlock()
				return // Exit the goroutine if the producer has finished
			}
			cond.Wait() // Wait for the signal from producer
		}
		fmt.Println("Consumed:", sharedData)
		sharedData = -1 // Reset shared data
		cond.L.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go producer(&wg)
	go consumer(&wg)

	wg.Wait()

	fmt.Println("Program execution completed.")
}

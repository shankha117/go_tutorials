package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int, 10) // Declare the channel inside main()

	wg.Add(20)

	for i := 0; i < 10; i++ {
		go sender(ch, 42+i)
	}

	go receiver(ch)

	wg.Wait()
}

func receiver(ch <-chan int) {

	// i := <-ch

	for i := range ch {

		fmt.Printf("RECEIVING %v \n", i)
		wg.Done()
	}
}

func sender(ch chan<- int, i int) {
	ch <- i
	wg.Done()
}

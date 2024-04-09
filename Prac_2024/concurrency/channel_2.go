package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		receiver(ch)
		wg.Done()
	}(ch)
	go func(ch chan<- int) {

		sender(ch, 42)
		sender(ch, 27)
		wg.Done()
		close(ch)

	}(ch)
	wg.Wait()
}

func receiver(ch <-chan int) {

	// i := <-ch
	// i_2 := <-ch

	// fmt.Printf("RECEIVING %v \n", i)
	// fmt.Printf("RECEIVING %v \n", i_2)

	// for i := range ch {

	// 	fmt.Printf("RECEIVING %v \n", i)

	// }

	for {
		if i, ok := <-ch; ok {
			fmt.Printf("RECEIVING %v \n", i)
		} else {
			fmt.Println(" channel closed")
			break

		}
	}

}

func sender(ch chan<- int, i int) {
	ch <- i
}

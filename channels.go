package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	wg.Add(4)
	go func(ch1, ch2, ch3 chan int, wg *sync.WaitGroup) {

		fmt.Println("GO message from ch 1", <-ch1)
		fmt.Println("GO message from ch 2", <-ch2)
		fmt.Println("GO message from ch 3", <-ch3)
		wg.Done()

	}(ch1, ch2, ch3, wg)

	// go func(ch chan int, wg *sync.WaitGroup) {

	// 	// fmt.Println("GO message from ch 2", <-ch)
	// 	wg.Done()

	// }(ch, wg)
	go func(ch1 chan int, wg *sync.WaitGroup) {
		val := 44
		ch1 <- val
		fmt.Println("ch1 _>", val)
		wg.Done()
	}(ch1, wg)
	go func(ch2 chan int, wg *sync.WaitGroup) {
		val := 49
		ch2 <- val
		fmt.Println("ch2 _>", val)
		wg.Done()
	}(ch2, wg)

	go func(ch3 chan int, wg *sync.WaitGroup) {
		val := 47
		ch3 <- val
		fmt.Println("ch3 _>", val)
		wg.Done()
	}(ch3, wg)

	wg.Wait()
}

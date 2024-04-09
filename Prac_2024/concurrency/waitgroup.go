package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

var counter = 0

// func main() {
// 	msg := "HELLO"

// 	wg.Add(1)

// 	go func(msg string) {
// 		fmt.Println(msg)
// 		wg.Done()
// 	}(msg)

// 	msg = "BYE"

// 	wg.Wait()

// }

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}

	wg.Wait()

}

func sayHello() {
	fmt.Printf("HELLO ## %v \n", counter)
	wg.Done()

}

func increment() {
	counter++
	wg.Done()
}

// func main() {
// 	hello := func(wg *sync.WaitGroup, id int) {
// 		defer wg.Done()
// 		fmt.Printf("Hello from %v!\n", id)
// 	}
// 	const numGreeters = 5
// 	var wg sync.WaitGroup
// 	wg.Add(numGreeters)
// 	for i := 0; i < numGreeters; i++ {
// 		go hello(&wg, i+1)
// 	}
// 	wg.Wait()
// }

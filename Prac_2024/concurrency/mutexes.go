package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

var counter = 0

var m = sync.RWMutex{}

func main() {

	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock() // ADD THE READ LOCK HERE
		go sayHello()
		m.Lock() // ADD THE WRITE LOCK HERE
		go increment()
	}

	wg.Wait()

}

func sayHello() {
	fmt.Printf("HELLO ## %v \n", counter)
	m.RUnlock() // READ UNLOCK
	wg.Done()

}

func increment() {
	counter++
	m.Unlock() // WRITE LOCK
	wg.Done()
}

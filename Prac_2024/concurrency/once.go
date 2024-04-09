package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var count = 0

func initialize() {
	count++
	fmt.Printf("Initializing... times %v \n", count)

}

func main() {
	// This will execute the initialize function only once
	once.Do(initialize)

	// Subsequent calls to Do will have no effect
	once.Do(initialize)

	// Subsequent calls to Do will have no effect
	once.Do(initialize)

	// Subsequent calls to Do will have no effect
	once.Do(initialize)

	fmt.Println("Program execution completed.")
}

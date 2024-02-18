package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		fmt.Println("Pluralsight")
	}()

}

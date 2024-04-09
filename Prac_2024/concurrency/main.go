package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "HELLO"

	// use arguments
	go func() {
		fmt.Println(msg)
	}()

	msg = "BYE"

	time.Sleep(100 * time.Millisecond)

}

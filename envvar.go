package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to go Env Variables ..")
	// fmt.Println(os.Environ())
	fmt.Println(os.Getenv("USERNAME"), "\\n")
	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }

}

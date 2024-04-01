package main

import (
	"fmt"
)

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func third() {
	fmt.Println("3rd")
}

func main() {
	defer second()
	defer third()
	defer first()

	a := "start"
	defer fmt.Println(a)
	a = "end"

}

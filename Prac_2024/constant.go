package main

import (
	"fmt"
)

const a = "abcd"

const i = iota

func main() {

	// const myConstat float64 = math.Sin(1.57)
	// fmt.Println(myConstat)
	fmt.Printf(" %v, %T \n ", a, a)

	const a = 123
	fmt.Printf(" %v, %T \n ", a, a)

	const e = 42
	// var c = 23
	var b int16 = 21

	fmt.Printf("%v , %T \n", e, e)
	fmt.Printf("%v , %T \n", e+b, e+b)
	// fmt.Println("%v , %T", c+b, c+b)

	// enumerated Constants

}

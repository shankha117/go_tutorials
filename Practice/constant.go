package main

import (
	"fmt"
)

// UPPER case first letter for a CONSTANT means we will be exporting that CONSTANT
// so , use the following
// when not trying to export --> myConstat
// when trying to export --> MyConstat

// CONSTANTS CAN BE SHADOWED !!!!!!

// const (
// 	a = iota
// 	b = iota
// 	c = iota
// )

const (
	_ = iota
	a = iota // 0
	b = iota
	c
	d = 200 * iota // this defines a new IOTA and the bellow values in the same block will follow the same pattern
	e
	f
)

const g = iota

const h int = iota

// var h int = iota
// # command-line-arguments
// Practice/constant.go:29:13: cannot use iota outside constant declaration

const (
	a2 = iota
)

func main() {

	// const a = 42
	// var b int16 = 27
	// fmt.Printf("%v  , %T\n", a+b, a+b)
	fmt.Printf("a -->> %v\n ", a)
	fmt.Printf("b -->> %v\n ", b)
	fmt.Printf("c -->> %v\n ", c)
	fmt.Printf("d -->> %v\n ", d)
	fmt.Printf("e -->> %v\n ", e)
	fmt.Printf("f -->> %v\n ", f)
	fmt.Printf("g -->> %v\n ", g)
	fmt.Printf("h -->> %v\n ", h)
	fmt.Printf("a2 -->> %v\n ", a2)

}

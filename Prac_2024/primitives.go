package main

import (
	"fmt"
	"reflect"
)

func Prims() {
	fmt.Println("THIS IS COOL")
}

func main() {

	var k bool = false

	fmt.Printf(" %v , %T \n", k, k)

	n := 1 == 1
	m := 2 == 2

	fmt.Printf("%v , %T\n", m, m)
	fmt.Printf("%v , %T\n", n, n)

	a, b := 10, 3

	// a -> 1010
	// b -> 0011

	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100

	s := "This is me"
	fmt.Println(s, reflect.TypeOf(s))
	fmt.Println(string(s[2]))
	fmt.Printf("%v , %T\n", s[2], s[2])

	// d := "This is me"
	// d[2] = 'k'

	sb := []byte(s)
	// list ascii values or the utf values in the string `s`
	fmt.Println(reflect.TypeOf(sb), sb, string(sb))

	// RUNES

	r := 'a'
	fmt.Printf(" %v, %T \n", r, r)

	rs := "a"
	fmt.Printf(" %v, %T \n", rs, rs)

	// var ru rune = "a"
	// fmt.Printf(" %v, %T \n", ru, ru)

}

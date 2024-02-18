package main

import (
	"fmt"
	"reflect"
)

func main() {

	var N int
	fmt.Printf("%v , %T\n", N, N)

	n := 1 == 1
	m := 2 == 2

	fmt.Printf("%v , %T\n", m, m)
	fmt.Printf("%v , %T\n", n, n)

	a := 12

	fmt.Println(a << 3)

	B := fmt.Sprintf("%b", a)

	fmt.Println("This is B ------->>>", B, reflect.TypeOf(B))

	s := "This is string"
	fmt.Println(s, reflect.TypeOf(s))
	fmt.Printf("%v , %T\n", s[2], s[2])
	// the reason s[2] is a uint8 -- the strings in go are actually an aliases for bytes .
	// strings are generally IMMUTABLE
	// # command-line-arguments
	// Practice/primitves.go:32:2: cannot assign to s[2] (value of type byte)
	// s[2] = "U"
	fmt.Printf("%v , %T\n", string(s[2]), string(s[2]))

	b := []byte(s)

	// list ascii values or the utf values in the string `s`
	fmt.Println(reflect.TypeOf(b), b)

	// RUNE RUNE RUNE
	// utf32

	// string type represent any utf8 char , but RUNE represents an UTF8 char

	r := 'a' // this is a RUNE as we are using SINGLE QUOTES

	fmt.Println(r, reflect.TypeOf(r))

	var R rune = 'a'

	fmt.Println(r, reflect.TypeOf(R))

}

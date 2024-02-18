package main

import (
	"fmt"
	"reflect"
)

var (
	name   = "shankha"
	course = "askans"
)
var aa = 3

func main() {
	// variable used inside the function uses this kind of syntax
	a := 10.233
	b := 2
	c := a + float64(b)
	fmt.Println("Name is ", name)
	fmt.Println("Course is ", course)
	fmt.Println(reflect.TypeOf(name), reflect.TypeOf(course))
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))
	fmt.Println(c)
}

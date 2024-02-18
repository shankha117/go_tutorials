package main

import (
	"fmt"
	"os"
	"reflect"
)

var (
	name   = "shankha"
	course = "askans"
)

func main() {

	a := os.Getenv("USERNAME")
	b := 2

	c := 32 + b
	fmt.Println("Name is ", a)
	fmt.Println("Course is ", course)

	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }

	fmt.Println(reflect.TypeOf(name), reflect.TypeOf(course))

	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))
	fmt.Println(c)
}

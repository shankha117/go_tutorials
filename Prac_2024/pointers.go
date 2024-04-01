package main

import (
	"fmt"
	"reflect"
)

type myStruct struct {
	foo int
}

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a, *b)

	fmt.Println(&a, b)

	*b = 53
	fmt.Println(a, *b)

	x := 1.5
	square(&x)
	fmt.Println(x)

	d := [3]int{1, 2, 3}
	e := &d[0]
	f := &d[1]
	fmt.Printf("%v %p %p\n", d, e, f) // [1 2 3] 0x1040a124 0x1040a128

	// struct
	var ms *myStruct
	ms = &myStruct{foo: 123}

	ms2 := myStruct{foo: 123}
	fmt.Println(ms)
	fmt.Println(&ms2)

	msNew := new(myStruct)
	fmt.Println(msNew)

	msNew.foo = 456
	fmt.Println(msNew)

	var msNil *myStruct
	fmt.Println(msNil)
	msNil = new(myStruct)
	fmt.Println(msNil)

}

func square(x *float64) {
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(*x))
	fmt.Printf("X value is %v", x)
	*x = *x * *x
}

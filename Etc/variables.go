package main

import (
	"fmt"
	"strconv"
)

var i int = 22

func main() {

	fmt.Println("HELLO")

	fmt.Println(i)

	var i int = 27

	fmt.Println(i)

	i = 42

	fmt.Println(i)

	fmt.Printf(" --- %v %T \n", i, i)

	var j string

	j = strconv.Itoa(i)

	fmt.Printf(" +++ %v %T \n", j, j)

}

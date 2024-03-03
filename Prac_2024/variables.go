package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var x = 1234
var inside string = "inside"

func getLength(email string) int {
	fmt.Println(inside)
	return len(email)
}

func main() {

	penniesPerText := 2
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)
	fmt.Println("this is the type of x", reflect.TypeOf(x))
	a := 12

	a_left_shifted_1 := a << 1

	fmt.Println("originla a is ", a)

	a_bhinary := fmt.Sprintf("%b", a_left_shifted_1)

	fmt.Println("Binary Rep of a is ", a_bhinary, a_left_shifted_1)
	fmt.Println(a_left_shifted_1)

	const email = "hahahhaha"

	fmt.Printf("%v \n\n", x)

	if length := getLength(email); length < 100 {
		fmt.Println("Email is invalid . length is ", length, email)
	}

	var i int = 42

	i = 1000

	fmt.Println(string(i))

	i = 42

	fmt.Println(i)

	fmt.Printf(" --- %v %T \n", i, i)

	var j string

	j = strconv.Itoa(i)

	fmt.Printf(" +++ %v %T \n", j, j)

	binary_int := 1100

	dec_str, _ := strconv.ParseInt(strconv.Itoa(binary_int), 2, 64)

	fmt.Println(dec_str, binary_int)

}

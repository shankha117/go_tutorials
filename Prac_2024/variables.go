package main

import (
	"fmt"
	"reflect"
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

	fmt.Println("Binary Rep of a is ", a_bhinary)
	fmt.Println(a_left_shifted_1)

	const email = "hahahhaha"

	fmt.Printf("%v \n\n", x)

	if length := getLength(email); length < 100 {
		fmt.Println("Email is invalid . length is ", length, email)
	}

	var i int = 42

	i = 3

	fmt.Println(i)

}

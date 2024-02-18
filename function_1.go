package main

import (
	"fmt"
	"reflect"
)

func passvalue(s1, s2 string) (str1, str2 string) {

	s1 = s1 + "jj"
	s2 = s2 + "jj"
	return s1, s2
}

func findmax(nums ...int) int {
	fmt.Println(nums)
	fmt.Println(reflect.TypeOf(nums))
	best := nums[0]

	for _, i := range nums {

		if best < i {
			best = i
		}

	}
	return best
}

func main() {
	fmt.Println("Welcome to Functions!")

	mod, auth := passvalue("shankha", "shuvro")

	fmt.Println(mod, auth)

	fmt.Println(findmax(4, 5, 6, 22, 9, 2, 6, 20, 911, 785415))

}

package main

import (
	"fmt"
)

func main() {
	x := [...]string{"Penn", "Teller"}
	numarray := [...]int{1, 2, 3, 4, 56, 9, 5, 2, 1}
	fmt.Println("Welcome to go Env Variables ..", x)
	fmt.Println("The num array is ", len(numarray), numarray)

	courses := []string{"shankha", "shuvro", "sinha"}

	for index, v := range courses {
		if index == 0 {
			fmt.Println("HI at index 0")
			continue
		}

		fmt.Println(v, index)

	}
}

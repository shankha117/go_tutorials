package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("Welcome to Slices!")

	course := []string{"shankha", "shuvro", "sinha"}

	fmt.Println(len(course), cap(course), course, reflect.TypeOf(course))

	fmt.Println(course[len(course):])

	arrsl := make([]int, 0, 4)
	fmt.Println("--", arrsl, len(arrsl), cap(arrsl))

	for i := 0; i < 17; i++ {
		arrsl = append(arrsl, i)
		fmt.Println(len(arrsl), cap(arrsl))
	}

	fmt.Println("This is arrsl", arrsl)

	newsl := []int{4, 5, 6, 8, 7}

	arrsl = append(arrsl, newsl...)
	fmt.Println("--", arrsl, len(arrsl), cap(arrsl))

	multiarray := [][]int{{1, 2, 4, 5}, {68, 7, 81, 2, 4, 814, 1, 2, 4, 5}}
	fmt.Println(multiarray)
	zeroarray := []uint{10}
	fmt.Println(zeroarray)
}

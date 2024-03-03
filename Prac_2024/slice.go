package main

import (
	"fmt"
	"strings"
)

func main() {

	s := []int{1, 2, 34}

	fmt.Printf(" Slice is %v and the length is %v  %T \n", s, len(s), s)

	a := []int{1, 2, 3}
	b := a
	a[2] = 455
	fmt.Println(a)
	fmt.Println(b)

	longSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(longSlice[:])
	fmt.Println(longSlice[2:])
	fmt.Println(longSlice[:4])
	fmt.Println(longSlice[2:5])

	// using MAKE
	x := make([]float64, 5, 10)
	fmt.Printf("Value is %v , Type is %T \n", x, x)

	// using Array Slices
	arr := [5]float64{1, 2, 3, 4, 5}
	arrSlice := arr[0:5]

	fmt.Printf("Value is %v , Type is %T \n", arr, arr)
	fmt.Printf("Value is %v , Type is %T \n", arrSlice, arrSlice)

	fmt.Printf("\n\n %v \n\n", strings.Repeat("=", 30))

	// slice1 := make([]int, 3, 3)

	slice1 := []int{1, 2, 3}

	fmt.Printf("Slice1 is %v", slice1)
	fmt.Printf(" Length of Slice1 %v , Cap is Slice1 %v \n", len(slice1), cap(slice1))

	fmt.Println("\nAPPENDING 4,5")
	slice2 := append(slice1, 4, 5)
	slice3 := append(slice2, 4, 5, 6, 7, 5, 6, 8, 9, 10, 16)

	fmt.Printf("\nSlice2 is %v", slice2)
	fmt.Printf(" Length of Slice2 %v , Cap is Slice2 %v\n", len(slice2), cap(slice2))

	fmt.Printf("\nSlice3 is %v", slice3)
	fmt.Printf(" Length of Slice3 %v , Cap is Slice3 %v\n", len(slice3), cap(slice3))

	fmt.Printf("\nSlice1 is %v", slice1)
	fmt.Printf(" Length of Slice1 %v , Cap is Slice1 %v\n", len(slice1), cap(slice1))

	fmt.Printf("\n\n %v \n\n", strings.Repeat("=", 30))

	// SPREAD OPERATOR

	sliceSpread := append(slice1, []int{3, 2, 1}...)
	fmt.Printf("SPREADDDD >>>>   %v \n", sliceSpread)

	// weirdness

	m := []int{1, 2, 3, 4, 5}
	fmt.Printf("m is %v , size is %v and cap is %v \n", m, len(m), cap(m))
	n := append(m[:2], m[3:]...)
	fmt.Printf("n is %v , size is %v and cap is %v\n", n, len(n), cap(n))
	fmt.Printf("m is %v , size is %v and cap is %v\n", m, len(m), cap(m))

	// instead what we can do is

	mm := []int{1, 2, 3, 4, 5}
	fmt.Printf(" This is mm %v \n", mm)
	nn := make([]int, 0, len(m)-1) // Create a new slice with length 0 and capacity len(m)-1
	nn = append(nn, mm[:2]...)     // Append elements from the beginning to index 2 of m
	nn = append(nn, mm[3:]...)     // Append elements from index 3 to the end of m
	fmt.Printf(" This is nn %v \n", nn)
	fmt.Printf(" This is mm %v \n", mm)

	// SLICE OF SLICE

	rowC := 5
	colC := 3

	matrix := make([][]int, 0)

	for i := 0; i < rowC; i++ {

		row := make([]int, 0, colC)

		for j := 0; j < colC; j++ {

			row = append(row, i*j)

		}

		matrix = append(matrix, row)

	}

	fmt.Println(matrix)

}

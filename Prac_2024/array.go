package main

import "fmt"

func main() {

	var IntArray [3]int
	var StrArray [2]string

	fmt.Printf(" int array ---> %v \n", IntArray)
	fmt.Printf(" string array ---> %v \n", StrArray)

	IntArray[0] = 1
	StrArray[1] = "shankha"

	fmt.Printf(" int array ---> %v \n", IntArray)
	fmt.Printf(" string array ---> %v \n", StrArray)

	fmt.Printf(" Length of int array ---> %v \n", len(IntArray))
	fmt.Printf(" Length of string array ---> %v \n", len(StrArray))

	grades := [3]int{1, 2, 3}
	fmt.Println(grades)

	gradesDyn := [...]int{1, 2, 3, 5, 6, 7}
	fmt.Println(gradesDyn)

	// multi dimentional array

	Identitiy := [3][3]int{[...]int{1, 2, 3}, [...]int{4, 5, 6}, [...]int{3, 5, 7}}
	fmt.Printf("\n Multi Dim ---> %v \n", Identitiy)

	// Arrays Are DEEP COPIED

	a := [...]int{1, 2, 3}
	b := a

	a[2] = 6

	fmt.Println(a)
	fmt.Println(b)

	// Add Pointers
	// b is going to point to the same data

	aOrg := [...]int{1, 2, 3}
	bPointer := &aOrg

	aOrg[2] = 6

	fmt.Println(aOrg)
	fmt.Println(bPointer)

}

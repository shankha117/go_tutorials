package main

import (
	"fmt"
)

func changeClass(class *int) int {
	fmt.Println("My class is ", *class)
	*class = 9
	fmt.Println("Now My class is ", *class)
	return *class
}

func main() {

	name := "shankha"
	class := 7
	isInSchool := true

	fmt.Println("Welcome to pointer !")
	fmt.Printf("\nName is %s and \nclass is %d and \nis he active -->%t\n", name, class, isInSchool)
	fmt.Println(" Mem address of class is ->", &class)
	fmt.Println("changing class by ref .. ")
	NewClass := &class
	fmt.Println(" Value of NewClass is ->", NewClass)
	fmt.Println(" Mem address of NewClass is ->", &NewClass)
	fmt.Println("Change my class !")
	class = 89
	newClassis := changeClass(&class)
	fmt.Println("My changed class is ", class)
	fmt.Println("My changed class is ", newClassis)

}

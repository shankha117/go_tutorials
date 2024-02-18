package main

import (
	"fmt"
)

func chnageCourse(course *string) string {

	*course = "new Value"
	fmt.Println("Course changd to ", *course)
	return *course
}

func main() {
	name := "shankha"
	course := "askans"

	fmt.Println(name, course)
	chnageCourse(&course)
	fmt.Println(name, course)

}

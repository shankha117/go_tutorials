package main_1

import (
	"fmt"
)

func chnageCourse(course *string) string {

	*course = "new Value"
	fmt.Println("Course changd to ", *course)
	return *course
}

func main_1() {
	name := "shankha"
	course := "askans"

	fmt.Println(name, course)
	chnageCourse(&course)
	fmt.Println(name, course)

}

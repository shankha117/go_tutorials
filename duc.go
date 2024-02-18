package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	name, course = "sjaks", "danbksd"
	myname       = "shankha"
)

func main() {
	moudle := 7
	fmt.Println("Welcome to Go")
	fmt.Printf("%s\n%s\n", name, course)
	fmt.Println("My Variable types are")
	fmt.Println("Name is", name, "type of name is", reflect.TypeOf(name))
	fmt.Println("Course is", course, "type of name is", reflect.TypeOf(course))
	fmt.Println("Moudle is", moudle, "type of name is", reflect.TypeOf(moudle))
	kaka := strconv.Itoa(moudle)
	fmt.Println(kaka)
	fmt.Println("kaka is", kaka, "type of name is", reflect.TypeOf(kaka))
	addstring := kaka + "_" + name
	fmt.Println("add_string is", addstring, "type of name is", reflect.TypeOf(addstring))

}

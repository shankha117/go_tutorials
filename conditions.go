package main

import "fmt"

func main() {
	fmt.Println("Welcome to Conditions!")

	if x := 5; x <= 10 {
		fmt.Println("X is less than or equals to 10")
	} else {
		fmt.Println("x is greater than 10")
	}

	x := "linux"

	switch x {
	case "linux":
		fmt.Println("linux courses")
	case "docker":
		fmt.Println("docker courses")

	case "windows":
		fmt.Println("windows courses")
	default:
		fmt.Println("No match")
	}

}

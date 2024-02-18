package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is Struct Example !")

	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	newvar := courseMeta{
		Author: "shankha",
		Level:  "Inter",
		Rating: 6,
	}

	fmt.Println(newvar.Author)

}

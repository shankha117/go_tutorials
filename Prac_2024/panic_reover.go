package main

import (
	"fmt"
	"log"
)

func main() {

	// Panic

	// b, c := 1, 0

	// ans := b / c

	// fmt.Println(ans)

	// fmt.Println("start")
	// defer fmt.Println(" DEFFERED ")
	// panic("some bad !")
	// fmt.Println("end")

	// fmt.Println("start")
	// defer func() {
	// 	x := 1 / 0
	// 	fmt.Println(x)
	// }()
	// fmt.Println("end")

	fmt.Println("start")
	panicker()
	fmt.Println("end")

}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("spmething bad happened")
	fmt.Println("done panicking")
}

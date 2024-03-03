package main

import "fmt"

func main() {

	// for i := 0; i < 10; i++ {

	// 	fmt.Println(i)

	// 	switch i {
	// 	case 2:
	// 		fmt.Println(" Printing 2 ")
	// 	case 4:
	// 		fmt.Println(" Printing 4 ")
	// 	case 6:
	// 		fmt.Println(" Printing 6 ")
	// 	case 19:
	// 		fmt.Println(" Printing 19 ")
	// 	}

	// }

	for i := 0; i < 10; i++ {

		fmt.Println(i)

		switch i {
		case 2, 3, 4:
			fmt.Println(" Printing 2 or 3 or 4 ")
		case 6, 71, 1:
			fmt.Println(" Printing 6 or 71 or 1 ")
		case 0, 9, 8:
			fmt.Println(" Printing 0 or 9 or 8 ")
		case 19:
			fmt.Println(" Printing 19 ")
		default:
			fmt.Printf(" ALWAYS PRINT %v \n", i)

		}

	}

}

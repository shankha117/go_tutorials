package main

import "fmt"

func main() {

	fmt.Println(" Foor loop")

	i := 1

	for i := 0; i < 3; i++ {

		fmt.Printf(" The value of i is --> %v \n", i)

	}

	fmt.Printf(" The i outside loop -->> %v\n", i)

	for i, j := 0, 0; i < 3; i, j = i+1, j+1 {

		fmt.Printf(" The value of i is %v j is %v\n", i, j)

	}

	k := 0

	for ; k < 3; k++ {
		fmt.Printf(" THIs is the value of K %v \n", k)
	}

	var j = 0
	for j < 5 {
		fmt.Printf(" THIS IS J %v \n", j)
		j++

	}

	for i := 1; ; i++ {

		if i == 10 {

			fmt.Println(" LIMIT REACHED ----")
			break
		}

	}

}

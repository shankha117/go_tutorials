package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func main() {

	// Convert from epoch to human readable date
	epoch := "1618603800001"[:10]

	int_epoch, _ := strconv.ParseInt(epoch, 10, 64)

	fmt.Printf("Epoch to human readable time is:\t%s\n", epochToHumanReadable(int_epoch))
	fmt.Println(epochToHumanReadable(int_epoch), reflect.TypeOf(epochToHumanReadable(int_epoch)))
	x := "10"
	fmt.Println(x, reflect.TypeOf(x))

	b1, _ := strconv.ParseFloat(x, 8)
	fmt.Println(b1, reflect.TypeOf(b1))

	numarray := [...]int{1, 2, 3, 4, 56, 9, 5, 2, 1}

	for i := 0; i < len(numarray); i++ {

		switch i {
		case 1, 2:
			fmt.Println(i, numarray[i])
		case 5, 6:
			fmt.Println(i, "--", numarray[i])
		}

	}

}

func epochToHumanReadable(epoch int64) string {
	return time.Unix(epoch, 0).String()
}

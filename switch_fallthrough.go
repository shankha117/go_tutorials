package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("Enven", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("ODD", tmpNum)
	}

	_, err := os.Open("/home/boombox/Desktop/command.txt")

	if err != nil {
		fmt.Println("Error -->", err)
	} else {
		fmt.Println("No error !")
	}

}

func random() int {
	rand.Seed(time.Now().Unix())
	return int(rand.Intn(10))

}

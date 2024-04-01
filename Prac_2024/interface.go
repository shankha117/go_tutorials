package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {

	fmt.Println("INTERFACE !!!")

	var w Write = ConsoleWriter{}

	fmt.Println(reflect.TypeOf(w))

	n, _ := w.Write([]byte("hello WORLD !"))

	fmt.Printf("%T , %v \n\n", n, n)

	// Assign the instances to the shape interface

	var s shape = circle{radius: 7}
	// var s shape = rectangle{length: 4, width: 3}

	// Type switch to determine the underlying type of s
	switch s.(type) {
	case circle:
		fmt.Println("s is a circle")
		fmt.Println("Area of the circle:", s.area())
	case rectangle:
		fmt.Println("s is a rectangle")
		fmt.Println("Area of the rectangle:", s.area())
	default:
		// Log an error if the underlying type is neither circle nor rectangle
		log.Fatal("Unknown shape")
	}

}

type Write interface {
	Write(sourceBytes []byte) (bytedcopied int, error error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {

	d := string(data)

	fmt.Printf(" THIS IS THE D --->>%v", d)

	n, err := fmt.Println(d)
	return n, err
}

// Define a shape interface
type shape interface {
	area() float64
}

// Define a circle struct
type circle struct {
	radius float64
}

// Implement the area method for the circle struct
func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// Define a rectangle struct
type rectangle struct {
	length float64
	width  float64
}

// Implement the area method for the rectangle struct
func (r rectangle) area() float64 {
	return r.length * r.width
}

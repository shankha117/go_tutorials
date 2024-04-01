package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

type Circle struct {
	x, y, r float64
}

type Doctor struct {
	num        int
	actorName  string
	companions []string
}

type car struct {
	Make   string
	Model  string
	Height int
	Width  int
}

type Animal struct {
	Name   string `required:"true" max:"100" default:"Tiger"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float64
	CanFly   bool
}

func getRandomName() string {
	names := []string{
		"Alice", "Bob", "Charlie", "David", "Emma", "Frank", "Grace", "Henry", "Ivy", "Jack",
	}

	// Create a local random generator with the current time as seed
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random index within the range of the names slice
	randomIndex := rng.Intn(len(names))

	// Return the name at the randomly selected index
	return names[randomIndex]
}

func main() {

	fmt.Println(" THIS IS STRUC !!!")

	docList := make([]Doctor, 0, 5)

	for i := 0; i < 5; i++ {

		newDoc := Doctor{
			num:       i,
			actorName: "shankha" + "_" + fmt.Sprint(i),
			companions: []string{
				getRandomName(),
				getRandomName(),
			},
		}

		docList = append(docList, newDoc)

	}

	docList[1].actorName = "NONONNONO"

	docList = append(docList, docList[0])

	// fmt.Println(docA, docA.companions[0])

	for _, v := range docList {

		fmt.Println(v)

	}

	// ANONYMOUS STRUCT
	myCar := struct {
		Make  string
		Model string
	}{Make: "tesla", Model: "model 3"}

	fmt.Println(myCar)

	newCar := car{
		Make:   "Tata",
		Model:  "Harrier",
		Height: 120,
		Width:  222,
	}

	fmt.Println(newCar)

	newCar2 := newCar

	newCar.Model = "Safari"

	fmt.Println(newCar)

	fmt.Println(newCar2)

	// NESTED STRUCTS

	c := Bird{}

	c.Name = "eagle"
	c.Origin = "India"
	c.SpeedKPH = 60.56
	c.CanFly = true

	// or use the Literal Syntax

	b := Bird{
		Animal:   Animal{Name: "EMU", Origin: "AUS"},
		SpeedKPH: 48,
		CanFly:   false,
	}

	fmt.Println(b)
	fmt.Println(c)

	animal := Animal{}

	// Get the type of the Animal struct
	t := reflect.TypeOf(animal)

	// Iterate over the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		// Get the field
		field := t.Field(i)

		// Get the tag value of the "required" tag
		requiredTag := field.Tag.Get("required")

		// Get the tag value of the "max" tag
		maxTag := field.Tag.Get("max")

		// Get the tag value of the "default" tag
		defaultTag := field.Tag.Get("default")

		// Print the field name and its tags
		fmt.Printf(strings.Repeat("=", 20) + "\n")
		fmt.Println("field:", field)
		fmt.Printf("required: %v , %T \n ", requiredTag, requiredTag)
		fmt.Println("max:", maxTag)
		fmt.Println("default:", defaultTag)

	}

	// Methods

	cir := Circle{0, 0, 5}
	fmt.Printf("Area of Circle %v \n", circleArea(cir))

	// pass the pointer
	fmt.Printf("Area of Circle %v \n", circleAreaP(&cir))

	// use the method
	fmt.Printf("Area of Circle %v \n", cir.area())

}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func circleAreaP(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

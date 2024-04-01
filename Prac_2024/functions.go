package main

import "fmt"

func main() {
	x := 5
	fmt.Println(x)
	increment(&x)

	fmt.Println(x)

	s := sum("THe sum is ", 1, 2, 3, 4, 5, 5)
	fmt.Println(s)

	ll := []int{1, 2, 3, 4, 4}

	ss := sum("NAA NANA", ll...)

	fmt.Printf("THIS IS SSSSS %v \n\n", ss)

	sNo := sumNoReturn("THe sum is ", 1, 2, 3, 4, 5, 5)
	fmt.Println(sNo)

	sp := sumPointer("THe sum is ", 1, 2, 3, 4, 5, 5)
	fmt.Println(*sp)

	d, err := divide(5, 2)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(d)

	// anonymous function
	func() {
		fmt.Println(" THIS IS Anonymous")
	}()

	// functions as variables

	f := func() {
		fmt.Println(" THIS IS VARIABLE NOW @")
	}

	f()

	var ff func() = func() {
		fmt.Println(" OKAY !")
	}

	ff()

	// Methods

	g := greeter{
		greeting: "HI",
		name:     "go",
	}

	g.greet()
	g.greet()

}

func increment(x *int) {
	*x++
}

func sum(msg string, values ...int) int {
	fmt.Println(values)

	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
	return result
}

func sumNoReturn(msg string, values ...int) (result int) {
	fmt.Println(values)

	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)

	return
}

func sumPointer(msg string, values ...int) *int {
	fmt.Println(values)

	result := 0

	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)

	return &result
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero ")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "NOOOO"
}

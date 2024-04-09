package main

import (
	"fmt"
	"sync"
)

// FactorialResult represents the result of factorial calculation
type FactorialResult struct {
	Number    int // Number for which factorial is calculated
	Factorial int // Resulting factorial value
}

func calculateFactorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * calculateFactorial(num-1)
}

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return &FactorialResult{}
		},
	}

	// Compute and store factorial results
	for i := 1; i <= 10; i++ {
		// Get a FactorialResult object from the pool or create a new one
		result := pool.Get().(*FactorialResult)

		// Set the number for which factorial is calculated
		result.Number = i

		// Reuse the factorial of the previous number if available in the pool
		if i > 1 {
			previousResult := pool.Get().(*FactorialResult)
			result.Factorial = previousResult.Factorial * i
			pool.Put(previousResult)
		} else {
			// Calculate the factorial for the first number
			result.Factorial = calculateFactorial(i)
		}

		// Print the factorial result
		fmt.Printf("Factorial of %d is %d\n", result.Number, result.Factorial)

		// Return the FactorialResult object to the pool for reuse
		pool.Put(result)
	}
}

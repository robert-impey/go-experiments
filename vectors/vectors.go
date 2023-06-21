package main

import (
	"fmt"
	"math"
)

// See Page 30 of "Linear Algebra for Dummies"

func main() {
	xs := []float64{3, 2, 4}

	sumOfSquares := 0.0

	for _, x := range xs {
		sumOfSquares += x * x
	}

	fmt.Println(math.Sqrt(sumOfSquares))
}

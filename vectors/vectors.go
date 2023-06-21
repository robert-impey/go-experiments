package main

import (
	"fmt"
	"math"
)

func main() {
	// See Page 30 of "Linear Algebra for Dummies"
	// Calculate the magnitude of a vector.
	xs := []float64{3, 2, 4}

	sumOfSquares := 0.0

	for _, x := range xs {
		sumOfSquares += x * x
	}

	mag := math.Sqrt(sumOfSquares)

	fmt.Println(mag)
}

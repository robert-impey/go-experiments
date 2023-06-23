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

	// Sum of Vectors (page 28)
	a := []float64{2, 4}
	b := []float64{3, 1}

	sumOfAAndB := make([]float64, len(a))

	for i, elementOfA := range a {
		elementOfB := b[i]
		sumOfAAndB[i] = elementOfA + elementOfB
	}

	fmt.Println("The sum of A and B")
	for _, elementOfSum := range sumOfAAndB {
		fmt.Println(elementOfSum)
	}

	// Dot product of two vectors
	us := []float64{4, -2, 0, 1}
	vs := []float64{-1, -3, 1, 5}

	dotProductItems := make([]float64, len(us))

	for i, u := range us {
		v := vs[i]
		dotProductItems[i] = u * v
	}

	fmt.Println("The dot product items of 2 vectors")
	for _, elem := range dotProductItems {
		fmt.Println(elem)
	}

	dotProduct := 0.0
	for _, elem := range dotProductItems {
		dotProduct += elem
	}

	fmt.Printf("The dot product: %f\n", dotProduct)
}

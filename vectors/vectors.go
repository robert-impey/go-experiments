package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// See Page 30 of "Linear Algebra for Dummies"
	// Calculate the magnitude of a vector.
	xs := []float64{3, 2, 4}
	fmt.Printf("xs = %s\n", vectorToString(xs))
	mag := calcMagnitude(xs)

	fmt.Printf("magnitude of xs = %f\n", mag)

	// Sum of Vectors (page 28)
	a := []float64{2, 4}
	b := []float64{3, 1}

	sumOfAAndB := make([]float64, len(a))

	for i, elementOfA := range a {
		elementOfB := b[i]
		sumOfAAndB[i] = elementOfA + elementOfB
	}

	fmt.Println("The sum of A and B")
	fmt.Println(vectorToString(sumOfAAndB))

	// Dot product of two vectors (page 37)
	us := []float64{4, -2, 0, 1}
	vs := []float64{-1, -3, 1, 5}

	dotProductItems := make([]float64, len(us))

	for i, u := range us {
		v := vs[i]
		dotProductItems[i] = u * v
	}

	fmt.Println("The dot product items of 2 vectors")
	fmt.Println(vectorToString(dotProductItems))

	dotProduct := 0.0
	for _, elem := range dotProductItems {
		dotProduct += elem
	}

	fmt.Printf("The dot product: %f\n", dotProduct)
}

func vectorToString(xs []float64) string {
	xStrs := make([]string, len(xs))
	for i, x := range xs {
		xStrs[i] = fmt.Sprintf("%f", x)
	}

	fmt.Print()

	return fmt.Sprintf("[%s]", strings.Join(xStrs, ", "))
}

func calcMagnitude(xs []float64) float64 {
	sumOfSquares := 0.0

	for _, x := range xs {
		sumOfSquares += x * x
	}

	return math.Sqrt(sumOfSquares)
}

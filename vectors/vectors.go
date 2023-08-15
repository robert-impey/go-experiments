package main

import (
	"errors"
	"fmt"
	"log"
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
	fmt.Printf("a = %s\n", vectorToString(a))
	fmt.Printf("b = %s\n", vectorToString(b))

	sumOfAAndB, err := calcSum(a, b)

	if err == nil {
		fmt.Printf("The sum of a and b %s\n", vectorToString(sumOfAAndB))
	} else {
		log.Fatal(fmt.Sprintf("Error: %s\n", err))
	}

	// Dot product of two vectors (page 37)
	u := []float64{4, -2, 0, 1}
	v := []float64{-1, -3, 1, 5}
	fmt.Printf("u = %s\n", vectorToString(u))
	fmt.Printf("v = %s\n", vectorToString(v))

	dotProduct, err := calcDotProduct(u, v)

	if err == nil {
		fmt.Printf("The dot product of u and v: %f\n", dotProduct)
	} else {
		log.Fatal(fmt.Sprintf("Error: %s\n", err))
	}
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

func calcSum(xs []float64, ys []float64) ([]float64, error) {
	if len(xs) != len(ys) {
		return nil, errors.New("mismatched lengths")
	}

	sumOfXsAndYs := make([]float64, len(xs))

	for i, x := range xs {
		y := ys[i]
		sumOfXsAndYs[i] = x + y
	}

	return sumOfXsAndYs, nil
}

func calcDotProduct(xs []float64, ys []float64) (float64, error) {
	if len(xs) != len(ys) {
		return 0, errors.New("mismatched lengths")
	}

	dotProduct := 0.0

	for i, x := range xs {
		y := ys[i]
		dotProduct += x * y
	}

	return dotProduct, nil
}

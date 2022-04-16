package main

import (
	"fmt"

	"github.com/robert-impey/go-experiments/golang-book/chapter11/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}

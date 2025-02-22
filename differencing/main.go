package main

import "fmt"

func main() {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, x := range data[1:] {
		fmt.Printf("data[i] = %f\n", data[i])
		fmt.Printf("x = %f\n", x)
		fmt.Printf("x-data[i] = %f\n", x-data[i])
	}
}

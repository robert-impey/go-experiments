package main

import "fmt"

func main() {
	var testNumbers = []int{1, 3, 5, 7, 10, 50, 100}

	for _, n := range testNumbers {
		findAndPrintSum(n)
	}
}

func findAndPrintSum(n int) {
	fmt.Printf("n = %d, sum = %d\n", n, findSum(n))
}

func findSum(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			sum += i
		}
	}

	return sum
}

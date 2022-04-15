package main

import "fmt"

func main() {
	lim := 30
	for i := 1; i <= lim; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizzbuzz!")
		} else if i%3 == 0 {
			fmt.Println("Fizz!")
		} else if i%5 == 0 {
			fmt.Println("Buzz!")
		} else {
			fmt.Println(i)
		}
	}
}

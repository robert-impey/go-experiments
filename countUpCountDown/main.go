package main

import "fmt"

var annPlace = 2.0
var benPlace = 4.0

var numSteps = 0

func main() {
	var annStep = 1.0 / 2
	var benStep = 1.0 / 8

	for ; annPlace <= benPlace; numSteps++ {
		printState()
		annPlace += annStep
		benPlace -= benStep
	}

	fmt.Println("They have passed each other!")
	
	printState()
}

func printState() {
	fmt.Printf("After %d steps, Ann is at %f and Ben is at %f\n", numSteps, annPlace, benPlace)
}

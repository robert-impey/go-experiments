package main

import (
	"flag"
	"fmt"
)

var p1Name = flag.String("p1", "Player 1", "Player 1 Name")
var p2Name = flag.String("p2", "Player 2", "Player 2 Name")

var p1Start = flag.Float64("p1st", 2.0, "Player 1 start")
var p2Start = flag.Float64("p2st", 4.0, "Player 2 Start")

var p1Step = flag.Float64("p1sp", 1.0/2, "Player 1 Step")
var p2Step = flag.Float64("p2sp", 1.0/8, "Player 2 Step")

var numSteps = 0

func main() {
	flag.Parse()

	var p1Place = *p1Start
	var p2Place = *p2Start

	for ; p1Place <= p2Place; numSteps++ {
		printState(&p1Place, &p2Place)
		p1Place += *p1Step
		p2Place -= *p2Step
	}

	fmt.Println("They have passed each other!")

	printState(&p1Place, &p2Place)
}

func printState(p1Place, p2Place *float64) {
	fmt.Printf("After %d steps, %s is at %f and %s is at %f\n",
		numSteps, *p1Name, *p1Place, *p2Name, *p2Place)
}

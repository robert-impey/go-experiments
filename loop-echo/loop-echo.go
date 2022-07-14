// Prints its command line arguments
package main

import (
	"fmt"
	"github.com/loov/hrtime"
	"os"
)

func main() {
	sep := " "
	var s string

	const numberOfExperiments = 4096

	bench := hrtime.NewBenchmark(numberOfExperiments)
	for bench.Next() {
		for _, arg := range os.Args[1:] {
			s += sep + arg
		}
	}

	fmt.Printf("String length %d\n", len(s))

	fmt.Printf("Loop echo times\n")
	fmt.Println(bench.Histogram(10))
}

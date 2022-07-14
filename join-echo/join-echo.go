// Prints its command line arguments
package main

import (
	"fmt"
	"github.com/loov/hrtime"
	"os"
	"strings"
)

func main() {
	sep := " "
	var s string

	const numberOfExperiments = 4096

	bench := hrtime.NewBenchmark(numberOfExperiments)
	for bench.Next() {
		s = strings.Join(os.Args[1:], sep)
	}

	fmt.Printf("String length %d\n", len(s))

	fmt.Printf("Join echo times\n")
	fmt.Println(bench.Histogram(10))
}

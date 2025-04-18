package main

import (
	"fmt"
	"os"
)

func main() {
	var envVars = []string{
		"HOME", // Works on Mac OSX
		"HOMEDIR",
		"USERPROFILE", // Works on Windows
	}

	for _, envVar := range envVars {
		if value, exists := os.LookupEnv(envVar); exists {
			fmt.Printf("%s: %s\n", envVar, value)
		}
	}
}

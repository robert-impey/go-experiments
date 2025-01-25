package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	str := id.String()
	fmt.Println(str)
}

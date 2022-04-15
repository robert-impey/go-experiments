// Server1 is a minimal "echo" server

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle) // each request calls handle
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handle
func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

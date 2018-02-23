package main

import (
	"fmt"
	"net/http"
	"os"
)

var calls = 0

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	calls++
	fmt.Fprintf(os.Stdout, "Request for: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Hello!\nCall number %d\n", calls)
}

func main() {
	http.HandleFunc("/", HelloWorld)

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

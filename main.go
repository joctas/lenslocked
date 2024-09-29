package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting server on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, world!</h1>")
}

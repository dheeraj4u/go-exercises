package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}

func header(w http.ResponseWriter, r *http.Request) {
	for name, header := range r.Header {
		for _, h := range header {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/header", header)

	http.ListenAndServe(":8081", nil)
}

// go run httpserver.go $
// curl localhost:8080/hello

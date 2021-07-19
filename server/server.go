package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
	for name, value := range req.Header {
		fmt.Println(name, value)
	}
}

func main() {
	http.HandleFunc("/", hello)
	address := ":8000"
	println("starting server on address" + address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}
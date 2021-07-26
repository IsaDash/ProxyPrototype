package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
	//for name, value := range req.Header {
	//	fmt.Println(name, value)
	//}

	// get client specific headers that were set in nginx proxy
	realIp := req.Header.Get("X-Real-IP")
	fowardedIp := req.Header.Get("X-Forwarded-For")
	verified := req.Header.Get("VERIFIED")
	dn := req.Header.Get("DN")

	fmt.Println(realIp)
	fmt.Println(fowardedIp)
	fmt.Println(verified)
	fmt.Println(dn)

	resp := `{"status": "CREATED"}`
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(resp))
	return
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
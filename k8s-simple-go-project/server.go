package main

import "net/http"

func main() {
	http.HandleFunc("/hello", SayHello)

	http.ListenAndServe(":8080", nil)
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World! Golang is awesome!"))
}

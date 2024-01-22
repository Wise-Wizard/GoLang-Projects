package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveHome)
	http.ListenAndServe(":9090", nil)
	fmt.Println("Server is listening on :9090...")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is the home page!")
}

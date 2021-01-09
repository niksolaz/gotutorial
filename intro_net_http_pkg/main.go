package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[1:] == "nicola" {
		fmt.Fprintf(w, "<div><h1>Hi %s</h1><p>I love Golang!</p></div>", r.URL.Path[1:])
		return
	}
	fmt.Fprintf(w, "<div><h1>Hi there,</h1><p>I love %s</p></div>", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

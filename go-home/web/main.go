package main

import (
	"io"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Local()
	io.WriteString(w, "Hello world! The current local time is " + now.String())
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
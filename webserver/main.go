package main

import (
	"io"
	"net/http"
)

func firstPage (w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "<h1>I'm skyling, this is my first page by go </h1>")
}

func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":2012", nil)
}
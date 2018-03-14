package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/qoute/1", p1)
	http.ListenAndServe(":8080", nil)
}

func p1(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, `<img src="q1.jpg">`)
}

package main

import (
	"io"
	"net/http"
)

type page1 int

func (p1 page1) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello page1 Hadle")
}

type page2 int

func (p2 page2) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello page2 handle")
}

func main() {
	var p1 page1
	var p2 page2
	http.Handle("/page1/", p1)
	http.Handle("/page2/", p2)

	http.ListenAndServe(":8080", nil)
}

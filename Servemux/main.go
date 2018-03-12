package main

import (
	"io"
	"net/http"
)

type p1 int

func (d p1) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "page 1")

}

type p2 int

func (b p2) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "page 2")

}
func main() {
	var d p1
	var b p2
	mux := http.NewServeMux()
	mux.Handle("/p1/", d)
	mux.Handle("/p2/", b)
	http.ListenAndServe(":8080", mux)
}

package main

import (
	"io"
	"net/http"
)

func p1(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "you're at page 1")

}

func p2(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "you're at page2")
}

func main() {

	http.HandleFunc("/p1/", p1)
	http.HandleFunc("/p2", p2)

	http.ListenAndServe(":9999", nil)
}

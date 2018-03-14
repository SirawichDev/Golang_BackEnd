package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/p2", p1)
	http.HandleFunc("/p3", p3)
	http.ListenAndServe(":8080", nil)
}

func p1(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, `<img src="q1.jpg">`)
}
func p3(res http.ResponseWriter, req *http.Request) {
	tpl
}

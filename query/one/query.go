package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("g")
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<form method="post"><input type="text" name="g">
		<input type="submit"></form><br>
	`+v)
}

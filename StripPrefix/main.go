package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", quote)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func quote(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res,"`<img src="/assets/tt.png">`)
}

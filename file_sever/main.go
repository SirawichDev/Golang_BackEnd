package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/p2", p1)
	http.HandleFunc("/p3", p3)
	http.ListenAndServe(":8080", nil)
}

func p1(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="q1.jpg">`)
}
func p3(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("404.html")

	if err != nil {
		log.Fatalln(err)
	}
	show := tpl.ExecuteTemplate(res, "404.html", nil)
	if show != nil {
		log.Fatalln(show)
	}
}

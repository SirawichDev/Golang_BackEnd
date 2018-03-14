package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func p1(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "you're at page 1")

}

func p2(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "you're at page2")
}
func p3(res http.ResponseWriter, req *http.Request) {
	fi, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	err = fi.ExecuteTemplate(res, "index.html", "20")
	if err != nil {
		fmt.Println(nil)
		log.Fatalln(err)

	}
}
func main() {

	http.HandleFunc("/p1/", p1)
	http.HandleFunc("/p2", p2)
	http.Handle("/p3/", http.HandlerFunc(p3))
	fmt.Println(nil)
	http.ListenAndServe(":9999", nil)

}

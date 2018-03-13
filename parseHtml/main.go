package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	var p1 page1
	var p2 page2
	mux := http.NewServeMux()
	mux.Handle("/1/", p1)
	mux.Handle("/2", p2)
	http.ListenAndServe(":8080", mux)
}

type page1 int

func (p1 page1) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "page one")
}

type page2 int

func (p2 page2) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	tm, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)

	}
	err = tm.ExecuteTemplate(res, "index.gohtml", "miew")
	if err != nil {
		log.Fatalln("error executing template", err)
	}

}

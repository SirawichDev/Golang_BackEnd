package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("show.html"))
}
func p2(res http.ResponseWriter, req *http.Request) {
	sage := []string{"hh", "vv", "ee", "bb"}
	err := tpl.ExecuteTemplate(res, "show.html", sage)
	if err != nil {
		log.Fatalln(err)
	}
}
func main() {

	http.Handle("/p2", http.HandlerFunc(p2))

	http.ListenAndServe(":8080", nil)

}

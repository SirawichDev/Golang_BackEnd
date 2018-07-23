package main

import (
	"html/template"
	"log"
	"net/http"
)

type Stu struct {
	Name string
	Code string
}

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("./templates/*.html"))
}
func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("name")
	p := req.FormValue("code")
	err := tmp.ExecuteTemplate(w, "index.html", Stu{v, p})
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("templates/*.html"))
}
func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, req *http.Request) {
	var s string
	if req.Method == http.MethodPost {
		fl, h, err := req.FormFile("files")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer fl.Close()
		fmt.Println("header :", h)
		bs, err := ioutil.ReadAll(fl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		s = string(bs)
	}
	err := tmp.ExecuteTemplate(w, "index.html", s)
	if err != nil {
		log.Fatalln(err)
	}

}

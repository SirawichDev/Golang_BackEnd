package main

import (
	"io"
	"net/http"
	
)
func main(){

	http.HandleFunc("/",quote)
	http.Handle("/assets/", http.StripPrefix("/assets",http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080",nil)
}

func quote(res http.ResponseWriter,req)
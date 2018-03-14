package main

import(
	"fmt"
	"net/http"
	"io"
	"log"
	"os"

)

func main(){
	http.HandleFunc("/",p1)
	http.HandleFunc("/quote.jpg",p2)

	http.ListenAndServe(":8080",nil)
	
}

func p2(res http.ResponseWriter,req *http.Request){
	f,err := os.Open("quote.jpg")
	if err !=nil{
		
		http.Errro(res,"File not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil{
		http.Error(res,"file not found",404)
		return
	}
		http.ServeContent(res , req, f.Name(), fi.ModTime(),f)
}
func p1(res http.ResponseWriter,req *http.Request){
	res.Header().Set("Content-Type","text/html; charset=utf-8")
	io.WriteString(res, '<img src ="quote.jpg">')
}
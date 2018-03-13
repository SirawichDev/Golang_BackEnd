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
		http.Error(w,"file not found",404)
		return
	}
	
}
func p1(res http.ResponseWriter,req *http.Request){
	g, err := os.
}
package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	index, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	new, err := os.Create("foo.html")
	if err != nil {
		log.Println("Error crating file", err)
	}
	defer new.Close()

	err = index.Execute(new, nil)
	if err != nil {
		log.Fatalln(err)
	}

}

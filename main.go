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

	err = index.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

}

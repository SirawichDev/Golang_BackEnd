package main

import (
	"fmt"
	"net/http"
)

type Server int

func (m Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hellow my serve miew miew")
}
func main() {
	var S Server
	http.ListenAndServe(":8080", S)

}

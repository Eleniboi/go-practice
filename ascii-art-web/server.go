package main

import (
	"fmt"
	"net/http"
)

//This  
func homeHandler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w, "Welcome Home")
}

func aboutHandler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w, "This page was created to teach young youth Tech")
}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":8080", nil)

}
package main

import (
	"ascii-web/logic"
	"fmt"
	"html/template"
	"net/http"
)

// homeHandler handle homepage requests
func asciiHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		fmt.Fprintln(w, "This route exists, but it is been access incorrectly")
		return
	}

	//the request from the web are been transfer into variables
	//one for the user input and the other for the banner choice

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	font, err := logic.LoadBanner("banners/" + banner)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	result := logic.GenerateArt(text, font)
	tmpl, err := template.ParseFiles("template/result.html")

	if err != nil {
		return
	}

	err = tmpl.Execute(w, result)

	if err != nil {
		return
	}

	//.fmt.Fprintln(w, result)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "<html><head></head><body><h1>how are you today</h1></body></html>")

}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)

	http.ListenAndServe(":8080", nil)
	fmt.Println("server is now live on port :8080..")

}

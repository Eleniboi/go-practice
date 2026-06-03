package main

import (
	"ascii-web/logic"
	"fmt"
	"net/http"
)

// homeHandler handle homepage requests
func asciiHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/ascii-art" {
		text := r.FormValue("text")
		banner := r.FormValue("banner")

		font, err := logic.LoadBanner(banner)

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		result := logic.GenerateArt(text, font)

		fmt.Fprintln(w, result)
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "template/index.html")
}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)

	http.ListenAndServe(":8080", nil)
	fmt.Println("server is now live on port :8080..")

}

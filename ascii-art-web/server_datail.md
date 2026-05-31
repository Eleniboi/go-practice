## About Go server and How it Works

**Handler** : is a function that run when a route is been visited, it is responsible for an incoming http request and constructing http response
```go
// homeHandler handle homepage requests
func homeHandler(w http.ResponseWriter, r *http.Request) {

	// route check, if the route is not the homepage("/") return a 404 error
	if r.URL.Path != "/" {
		fmt.Fprintln(w, "404 Not Found ! ! !")
		return
	}

	fmt.Fprintln(w, "How fa na bro, it been a long time,\ni thought you are no longer in otukpo")
}

func main() {

	// HandleFunc is a built in function that registers route using the handler function 
	http.HandleFunc("/", homeHandler)
	// ListenAndServe keeps the go server alive 24/7, waiting for a server to connect 
	http.ListenAndServe(":8080", nil)
}
```
- NOTE: more than one route can run at a time simultane

## HTML Template
- A template is use to seprate logic from design.

the above server send plain text to the browser,

Now the text need to be on a actual webpage, 

```go
// It works this way but it make the go server messy and hard to read 
// so that is why we need the template to separate logic from design
func homeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "<html><head></head><body><h1>how are you today</h1></body></html>")

}
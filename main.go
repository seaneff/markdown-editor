// to run locally:
// go build
// ./fileserver

// this is the starting point of the web app
package main 

// import net/http package
// and package to format markdown

import (
	"net/http"
	"github.com/russross/blackfriday"
)

// main function
func main () {

	// fileserver code

	// use handlefunc and handle methods to define routing
	// calling the "/" pattern acts as a catch-all route (defined last)
	http.HandleFunc("/markdown", generateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))

	// start the server
	// serve the files in the current directory on the port 8080
	// the handler is nil, assumes HTTP requests will be handled by the
	// net/http packages default (it's called http.ServeMux)
	http.ListenAndServe(":8080", nil)

}

// generateMarkdown function implements handlerfunc interface
// and renders HTML from a form field that contains markdown
func generateMarkdown(rw http.ResponseWriter, r *http.Request) {
	// content is retrieved with r.FormValue("body")
	// similar examples are r.Header, r.URL, etc
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	// finalize response
	rw.Write(markdown)
}
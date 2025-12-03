package main

// Main seems to be the entry point of most go apps. similar to index.js in Node.js. I'm still a
// bit unsure on best practices for structure. I've made a second package as form but I'm assuming this
// will be installable on its own where this isn't intended, and I have read that I could potentially use an internal???
// Probably could also seperate out logic too. 

import (
	"log"
	"net/http"

	"github.com/matthewparkin/go-api/form"
)

func main() {
	//    This handles the core functionality: retrieving form definitions.
	svc := form.NewService()

	//    HTTP handler - transport layer.
	handler := form.NewHTTPHandler(svc)

	//   Start the HTTP server on port 8080 and listen.
	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", handler)
}

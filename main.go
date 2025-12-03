package main

import (
	"log"
	"net/http"

	"github.com/matthewparkin/go-api/form"
)

func main() {
	svc := form.NewService()
	handler := form.NewHTTPHandler(svc)
	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", handler)
}

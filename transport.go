package main

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
)

func decodeHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req Req
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {
	svc := ServiceImplementation{}
	helloHandler := kithttp.NewServer(
		MakeEndpoint(svc),
		decodeHelloRequest,
		encodeResponse,
	)

	http.Handle("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

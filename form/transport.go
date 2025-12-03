package form

import (
	"context"
	"encoding/json"
	"net/http"

	gohttp "github.com/go-kit/kit/transport/http"
)

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	formID := r.URL.Query().Get("id")
	return GetRequest{FormID: formID}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func NewHTTPHandler(service Service) http.Handler {
	formHandler := gohttp.NewServer(
		MakeEndpoint(service),
		decodeRequest,
		encodeResponse,
	)
	mux := http.NewServeMux()
	mux.Handle("/form", formHandler)
	return mux
}

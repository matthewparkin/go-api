package form

import (
	"context"
	"encoding/json"
	"net/http"

	gohttp "github.com/go-kit/kit/transport/http"
)

// DecodeRequest extracts form ID from query param "id".
func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	formID := r.URL.Query().Get("id")
	return GetRequest{FormID: formID}, nil
}

// encodeResponse writes response as JSON.
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// NewHTTPHandler wires service -> endpoint -> HTTP handlers. Still a bit shaky on this, need to look into it more. 
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

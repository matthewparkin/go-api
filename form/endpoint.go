package form

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// MakeEndpoint wraps Service in Go-Kit endpoint pattern.
// I'm getting goo linting warning for interface saying I  can any this?
func MakeEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		elements, err := service.GetForm(req.FormID)
		if err != nil {
			return nil, err
		}
		return GetResponse{Elements: elements}, nil
	}
}

// GetRequest: form ID from query param.
type GetRequest struct {
	FormID string `json:"form_id"`
}

// GetResponse: form elements.
type GetResponse struct {
	Elements []FormElement `json:"elements"`
}

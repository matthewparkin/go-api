package form

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		elements, err := service.GetForm(req.FormID)
		if err != nil {
			return GetResponse{Elements: nil, Err: err.Error()}, nil
		}
		return GetResponse{Elements: elements, Err: ""}, nil
	}
}

type GetRequest struct {
	FormID string `json:"form_id"`
}

type GetResponse struct {
	Elements []FormElement `json:"elements"`
	Err      string        `json:"err,omitempty"`
}

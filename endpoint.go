package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequest)
		elements, err := service.GetForm(req.FormID)
		if err != nil {
			return getResponse{Elements: nil, Err: err.Error()}, nil
		}
		return getResponse{Elements: elements, Err: ""}, nil
	}
}

type getRequest struct {
	FormID string `json:"form_id"`
}

type getResponse struct {
	Elements []FormElement `json:"elements"`
	Err      string        `json:"err,omitempty"`
}

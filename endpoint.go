package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Req struct {
	Name string `json:"name"`
}

type Res struct {
	Message string `json:"message"`
}

func MakeEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Req)
		msg := svc.SayHello(req.Name)
		return Res{Message: msg}, nil
	}
}

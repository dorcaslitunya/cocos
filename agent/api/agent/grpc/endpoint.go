package grpc

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/ultravioletrs/agent/agent"
)

func runEndpoint(svc agent.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(runReq)

		if err := req.validate(); err != nil {
			return runRes{}, err
		}

		computation, err := svc.Run(context.TODO(), req.computation)
		if err != nil {
			return runRes{}, err
		}

		return runRes{Computation: computation}, nil
	}
}

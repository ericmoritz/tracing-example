package backend

import (
	"fmt"

	otgrpc "github.com/opentracing-contrib/go-grpc"
	opentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

func NewClient() (HelloServiceClient, error) {
	tracer := opentracing.GlobalTracer()
	// Set up a connection to the server.
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(tracer)),
		grpc.WithStreamInterceptor(
			otgrpc.OpenTracingStreamClientInterceptor(tracer)),
	)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to service: %v", err)
	}
	return NewHelloServiceClient(conn), nil
}

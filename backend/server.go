package backend

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	tracerPkg "tracing-example/tracer"

	otgrpc "github.com/opentracing-contrib/go-grpc"
	"google.golang.org/grpc"
)

const (
	addr = "0.0.0.0:19090"
)

func Main(wg *sync.WaitGroup) {
	defer wg.Done()
	tracer, closer, err := tracerPkg.New("tracing-example-backend")
	if err != nil {
		log.Fatalf("ERROR: cannot init tracer: %v\n", err)
	}

	defer closer.Close()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(tracer)),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(tracer)),
	)
	RegisterHelloServiceServer(s, &server{})
	log.Printf("backend listening on %v", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error listening: %v", err)
	}
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	out := &HelloResponse{}
	out.Reply = fmt.Sprintf("Hello, %s", in.Name)
	return out, nil
}

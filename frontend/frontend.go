package frontend

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"tracing-example/backend"
	"tracing-example/tracer"

	opentracing "github.com/opentracing/opentracing-go"
)

const (
	addr = "0.0.0.0:18080"
)

func Main(wg *sync.WaitGroup) {
	defer wg.Done()
	tracer, closer, err := tracer.New("tracing-example-frontend")
	if err != nil {
		log.Fatalf("ERROR: cannot init tracer: %v\n", err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	backendClient, err := backend.NewClient()
	if err != nil {
		log.Fatalf("Error creating client to backend: %v", err)
	}
	mux := http.NewServeMux()
	f := &frontend{backendClient}
	mux.Handle("/hello-world", http.HandlerFunc(f.HelloWorld))
	log.Printf("HTTP frontend listening on %v", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

type frontend struct {
	backendClient backend.HelloServiceClient
}

func (f *frontend) HelloWorld(w http.ResponseWriter, r *http.Request) {
	var span opentracing.Span
	span = opentracing.GlobalTracer().StartSpan("HelloWorld")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	resp, err := f.backendClient.SayHello(ctx, &backend.HelloRequest{Name: "World"})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error talking to backend: %v", err), 500)
		return
	}
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprint(w, resp.Reply)
}

package tracer

import (
	"io"
	"log"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

func New(name string) (opentracing.Tracer, io.Closer, error) {
	cfg, err := config.FromEnv()
	if err != nil {
		return nil, nil, err
	}
	log.Printf("DEBUG: tracing cfg.Reporter: %#v\n", cfg.Reporter)
	log.Printf("DEBUG: tracing cfg.Sampler: %#v\n", cfg.Sampler)

	return cfg.New(name, config.Logger(jaeger.StdLogger))
}

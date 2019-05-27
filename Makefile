all: test build

test: backend/backend.pb.go
	go test .

build: backend/backend.pb.go
	go build .

backend/backend.pb.go:
	go get google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go
	protoc -I backend backend/backend.proto --go_out=plugins=grpc:backend

.PHONY: all

run: test build
	go run .

helm:
	docker build -t tracing-example:`cat VERSION` .

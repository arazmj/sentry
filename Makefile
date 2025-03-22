.PHONY: proto build run

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		api/proto/sentry.proto

build:
	go build -o bin/sentry-server server/main.go
	go build -o bin/sentry cmd/cli/main.go

run: build
	./bin/sentry-server

deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest 
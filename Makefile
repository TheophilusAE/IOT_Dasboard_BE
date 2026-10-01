.PHONY: run build test vet verify-ws

run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

test:
	go test ./...

vet:
	go vet ./...

verify-ws:
	go run ./scripts/verify/ws_listen.go

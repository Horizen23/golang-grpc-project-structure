
install:
	go mod download
wire:
	cd cmd/server/grpc && wire

gen-protos:
	scripts/generate_protos.sh

start:
	go run ./cmd/server/main.go

build:
	go build -ldflags="-s -w" -o ./app ./cmd/server/main.go

dev-server:
	air -c air.server.toml

dev-client:
	air -c air.client.toml

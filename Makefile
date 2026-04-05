run:
	go run ./cmd/main.go

build:
	go build -o ./bin/cacheserver ./cmd/main.go

test:
	go test ./...

test-race:
	go test -race ./...

.PHONY: run build test test-race

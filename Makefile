run:
	go run ./cmd/main.go

build:
	go build -o ./bin/cacheserver ./cmd/main.go

test:
	go test ./...

.PHONY: run build test

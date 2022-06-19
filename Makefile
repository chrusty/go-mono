test:
	@go test ./...

build:
	@mkdir -p bin
	@go build -o bin/go-mono main.go

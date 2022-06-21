test:
	@go test ./... -cover

build:
	@mkdir -p bin
	@go build -o bin/go-mono main.go

install:
	@go install

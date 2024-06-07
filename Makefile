build:
	@go build -o bin/cixac-interpreter cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/cixac-interpreter
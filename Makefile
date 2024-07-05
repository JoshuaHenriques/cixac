build:
	@go build -o bin/cixac cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/cixac
